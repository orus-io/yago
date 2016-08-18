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

var magicYagormComment = regexp.MustCompile(`yagorm:([0-9A-Za-z_\.]+)?`)

func guessColumnType(goType string) string {
	if goType == "int64" {
		return "qb.BigInt()"
	} else if goType == "string" {
		return "qb.Varchar().NotNull()"
	} else if goType == "*string" {
		return "qb.Varchar()"
	} else if goType == "time.Time" {
		return "qb.Timestamp().NotNull()"
	} else if goType == "*time.Time" {
		return "qb.Timestamp()"
	}
	panic(fmt.Sprintf("Cannot guess column type for go type %s", goType))
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
		panic(fmt.Sprintf("yagorm: getGoType: unhandled '%s' (%#v). Please report this bug.", x, x))
	}
}

func parseStructTypeSpecs(ts *ast.TypeSpec, str *ast.StructType) (*StructData, error) {
	res := &StructData{
		Name:   ts.Name.Name,
		Fields: nil,
	}

	for _, f := range str.Fields.List {
		tag := ""
		if f.Tag != nil && len(f.Tag.Value) > 2 {
			tag = f.Tag.Value
			tag = reflect.StructTag(tag[1 : len(tag)-1]).Get("yagorm")
		}
		if len(f.Names) == 0 {
			if tag == "" {
				continue
			}
			return nil, fmt.Errorf(`yagorm: %s has anonymous field %s with "yagorm:" tag, it is not allowed`, res.Name, f.Type)
		}
		if len(f.Names) != 1 {
			panic(fmt.Sprintf("yagorm: %d names: %#v. Don't know what to do.", len(f.Names), f.Names))
		}

		name := f.Names[0]

		if !name.IsExported() {
			if tag == "" {
				continue
			}
			return nil, fmt.Errorf(`yagorm: %s has non-exported field %s with "yagorm:" tag, it is not allowed`, res.Name, name)
		}

		goType := getGoType(f.Type)

		res.Fields = append(res.Fields, FieldData{
			Name:       name.Name,
			Type:       goType,
			ColumnName: name.Name,
			ColumnType: guessColumnType(goType),
		})
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

			sm := magicYagormComment.FindStringSubmatch(doc.Text())
			if len(sm) == 0 {
				continue
			}

			tablename := strings.ToLower(ts.Name.Name)
			if len(sm) > 1 && sm[1] != "" {
				tablename = sm[1]
			}

			str, ok := ts.Type.(*ast.StructType)
			if !ok || str.Incomplete {
				continue
			}

			sd, err := parseStructTypeSpecs(ts, str)
			if err != nil {
				return nil, err
			}
			sd.TableName = tablename
			res = append(res, sd)
		}
	}

	return res, nil
}
