package generate

// Big chunks of code here are initialy copied from reform
// (see https://github.com/go-reform/reform/blob/v1-stable/parse)

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"regexp"
	"strings"
)

var magicYagoComment = regexp.MustCompile(`yago:([0-9A-Za-z_\.,]+)?`)

type structDefArgs struct {
	TableName string
	AutoAttrs bool
}

func readNameValue(s string) (name string, value string) {
	l := strings.SplitN(s, "=", 2)
	name = l[0]
	value = l[1]
	return
}

func magicYagoCommentArgs(doc string) (args structDefArgs, ok bool) {
	sm := magicYagoComment.FindStringSubmatch(doc)
	if len(sm) == 0 {
		return
	}

	if len(sm) > 1 && sm[1] != "" {
		splitted := strings.Split(sm[1], ",")
		for _, arg := range splitted {
			if arg == "autoattrs" {
				args.AutoAttrs = true
			} else {
				args.TableName = arg
			}
		}
	}

	ok = true

	return
}

func readColumnTags(tag string) (tags ColumnTags) {
	splitted := strings.Split(tag, ",")
	for _, arg := range splitted {
		if strings.Index(arg, "=") != -1 {
			name, value := readNameValue(arg)
			if name == "index" {
				tags.Indexes = append(tags.Indexes, value)
			} else if name == "unique_index" {
				tags.UniqueIndexes = append(tags.UniqueIndexes, value)
			} else {
				panic(fmt.Sprintf("Invalid tag %v", arg))
			}
		} else if arg == "primary_key" {
			tags.PrimaryKey = true
		} else if arg == "auto_increment" {
			tags.AutoIncrement = true
		} else if arg == "index" {
			tags.Indexes = append(tags.Indexes, ".")
		} else if arg == "unique_index" {
			tags.UniqueIndexes = append(tags.UniqueIndexes, ".")
		} else if arg == "." {
		} else {
			panic(fmt.Sprintf("Unknown tag %v", arg))
		}
	}
	return
}

func getGoType(x ast.Expr) string {
	switch t := x.(type) {
	case *ast.StarExpr:
		return "*" + getGoType(t.X)
	case *ast.Ident:
		return t.String()
	case *ast.SelectorExpr:
		return getGoType(t.X) + "." + getGoType(t.Sel)
	default:
		panic(fmt.Sprintf("yago: getGoType: unhandled '%s' (%#v). Please report this bug.", x, x))
	}
}

func parseStructTypeSpecs(ts *ast.TypeSpec, str *ast.StructType, autoattrs bool) (*StructData, error) {
	res := &StructData{
		Name:          ts.Name.Name,
		Fields:        nil,
		Indexes:       make(map[string][]int),
		UniqueIndexes: make(map[string][]int),
	}

	for _, f := range str.Fields.List {
		hasTags := false
		tags := ColumnTags{}
		if f.Tag != nil && len(f.Tag.Value) > 2 {
			tag := f.Tag.Value
			tag = reflect.StructTag(tag[1 : len(tag)-1]).Get("yago")
			if tag != "" {
				hasTags = true
				tags = readColumnTags(tag)
			}
		}
		if len(f.Names) == 0 {
			if hasTags {
				return nil, fmt.Errorf(
					`yago: %s has anonymous field %s with "yago:" tag, it is not allowed`,
					res.Name, f.Type)
			}
			continue
		}
		if len(f.Names) != 1 {
			panic(fmt.Sprintf("yago: %d names: %#v. Don't know what to do.", len(f.Names), f.Names))
		}

		name := f.Names[0]

		if hasTags && !name.IsExported() {
			return nil, fmt.Errorf(`yago: %s has non-exported field %s with "yago:" tag, it is not allowed`, res.Name, name)
		}
		if !(hasTags || autoattrs && name.IsExported()) {
			continue
		}

		goType := getGoType(f.Type)

		field := FieldData{
			Tags: tags,
			Name: name.Name,
			Type: goType,
		}
		res.Fields = append(res.Fields, field)
		fieldIndex := len(res.Fields) - 1

		for _, index := range field.Tags.Indexes {
			if _, ok := res.Indexes[index]; ok {
				res.Indexes[index] = append(res.Indexes[index], fieldIndex)
			} else {
				res.Indexes[index] = []int{fieldIndex}
			}
		}
		for _, index := range field.Tags.UniqueIndexes {
			if _, ok := res.UniqueIndexes[index]; ok {
				res.UniqueIndexes[index] = append(res.UniqueIndexes[index], fieldIndex)
			} else {
				res.UniqueIndexes[index] = []int{fieldIndex}
			}
		}
	}

	return res, nil
}

// ParseFile parses a file and returns found structs that should be mapped
func ParseFile(path string) ([]*StructData, error) {
	fset := token.NewFileSet()
	fileNode, err := parser.ParseFile(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("Error parsing file %v: %v", path, err)
	}

	var res []*StructData

	for _, decl := range fileNode.Decls {
		gd, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range gd.Specs {
			ts, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			// magic comment may be attached to "type Foo struct" (TypeSpec)
			// or to "type (" (GenDecl)
			doc := ts.Doc
			if doc == nil && len(gd.Specs) == 1 {
				doc = gd.Doc
			}
			if doc == nil {
				continue
			}

			args, ok := magicYagoCommentArgs(doc.Text())

			if !ok {
				continue
			}

			tablename := strings.ToLower(ts.Name.Name)

			if args.TableName != "" {
				tablename = args.TableName
			}

			str, ok := ts.Type.(*ast.StructType)
			if !ok || str.Incomplete {
				continue
			}

			sd, err := parseStructTypeSpecs(ts, str, args.AutoAttrs)
			if err != nil {
				return nil, err
			}
			sd.TableName = tablename
			res = append(res, sd)
		}
	}

	return res, nil
}
