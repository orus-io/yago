package generate

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func guessColumnType(goType string) string {
	if goType == "int64" {
		return "qb.BigInt()"
	} else if goType == "string" {
		return "qb.Varchar()"
	} else if goType == "*string" {
		return "qb.Varchar()"
	} else if goType == "time.Time" {
		return "qb.Timestamp()"
	} else if goType == "*time.Time" {
		return "qb.Timestamp()"
	} else if goType == "uuid.UUID" {
		return "qb.UUID()"
	}
	panic(fmt.Sprintf("Cannot guess column type for go type %s", goType))
}

func makeColumnName(name string) string {
	return ToDBName(name)
}

func getEmptyValue(goType string) string {
	if goType[0] == '*' {
		return "nil"
	} else if goType == "string" {
		return `""`
	} else if goType[0:3] == "int" {
		return "0"
	} else if goType == "time.Time" {
		return "(time.Time{})"
	} else if goType == "uuid.UUID" {
		return "(uuid.UUID{})"
	}
	panic(fmt.Sprintf("I have no empty value for type '%v'", goType))
}

func prepareFieldData(str *StructData, f *FieldData) {
	if f.ColumnName == "" {
		f.ColumnName = makeColumnName(f.Name)
	}
	if f.ColumnType == "" {
		f.ColumnType = guessColumnType(f.Type)
	}
	if f.ColumnModifiers == "" {
		if f.Tags.PrimaryKey {
			f.ColumnModifiers += ".PrimaryKey()"
		}
		if f.Tags.AutoIncrement {
			f.ColumnModifiers += ".AutoIncrement()"
			str.AutoIncrementPKey = f
		}
		if f.Type[0] == '*' {
			f.ColumnModifiers += ".Null()"
		} else {
			f.ColumnModifiers += ".NotNull()"
		}
	}
	if f.EmptyValue == "" {
		f.EmptyValue = getEmptyValue(f.Type)
	}
	if f.NameConst == "" {
		f.NameConst = fmt.Sprintf(
			"%s%s",
			str.Name, f.Name,
		)
	}
	if f.ColumnNameConst == "" {
		f.ColumnNameConst = fmt.Sprintf(
			"%s%sColumnName",
			str.Name, f.Name,
		)
	}
}

func prepareStructData(str *StructData, fd FileData) {
	str.File = fd
	str.PrivateBasename = strings.ToLower(str.Name[0:1]) + str.Name[1:]
	for i := range str.Fields {
		prepareFieldData(str, &str.Fields[i])
	}
}

func postPrepare(filedata *FileData, structs map[string]*StructData) {
	for _, str := range structs {
		for _, name := range str.Embed {
			embedded, ok := structs[name]
			if ok {
				for index, fields := range embedded.Indexes {
					if _, ok := str.Indexes[index]; !ok {
						str.Indexes[index] = []int{}
					}
					for _, fieldIndex := range fields {
						str.Indexes[index] = append(str.Indexes[index], len(str.Fields)+fieldIndex)
					}
				}
				for _, field := range embedded.Fields {
					field.FromEmbedded = true
					str.Fields = append(str.Fields, field)
				}
			}
		}
		for i := range str.Fields {
			if str.Fields[i].Tags.PrimaryKey {
				str.PKeyFields = append(str.PKeyFields, &str.Fields[i])
			}
		}
	}
	for _, str := range structs {
		for i, f := range str.Fields {
			if f.Tags.PrimaryKey && str.Fields[i].Type == "uuid.UUID" {
				filedata.Imports["github.com/m4rw3r/uuid"] = true
			}
			for _, fkDef := range str.Fields[i].Tags.ForeignKeys {
				var (
					fk           string
					structName   string
					refFieldName string
					refStruct    *StructData
					refField     *FieldData
					onUpdate     string
					onDelete     string
				)
				if strings.Index(fkDef, " ") != -1 {
					splitted := strings.Split(fkDef, " ")
					fk = splitted[0]
					for i, w := range splitted[1 : len(splitted)-1] {
						fmt.Println(w, splitted[i+2])
						if strings.ToUpper(w) == "ONUPDATE" {
							onUpdate = strings.ToUpper(splitted[i+2])
						}
						if strings.ToUpper(w) == "ONDELETE" {
							onDelete = strings.ToUpper(splitted[i+2])
						}
					}
				} else {
					fk = fkDef
				}
				if strings.Index(fk, ".") != -1 {
					splitted := strings.Split(fk, ".")
					structName = splitted[0]
					refFieldName = splitted[1]
				} else {
					structName = fk
				}
				refStruct = structs[structName]
				if refFieldName == "" {
					refField = refStruct.PKeyFields[0]
				} else {
					for i := range refStruct.Fields {
						if refStruct.Fields[i].Name == refFieldName {
							refField = &refStruct.Fields[i]
						}
					}
				}

				str.ForeignKeys = append(str.ForeignKeys, FKData{
					Column:    &str.Fields[i],
					RefTable:  refStruct,
					RefColumn: refField,
					OnUpdate:  onUpdate,
					OnDelete:  onDelete,
				})
			}
		}
	}
}

// ProcessFile processes a go file and generates mapper and mappedstruct
// interfaces implementations for the yago structs.
func ProcessFile(logger *log.Logger, path string, file string, pack string, output string) error {

	ext := filepath.Ext(file)
	base := strings.TrimSuffix(file, ext)

	if output == "" {
		output = filepath.Join(path, base+"_yago"+ext)
	}

	filedata := FileData{Package: pack, Imports: make(map[string]bool)}

	structs, err := ParseFile(filepath.Join(path, file))
	if err != nil {
		return err
	}

	structsByName := make(map[string]*StructData)
	for _, str := range structs {
		prepareStructData(str, filedata)
		structsByName[str.Name] = str
	}
	postPrepare(&filedata, structsByName)

	outf, err := os.Create(output)
	if err != nil {
		return err
	}

	defer outf.Close()

	if err := prologTemplate.Execute(outf, &filedata); err != nil {
		return err
	}

	for _, str := range structs {
		if err := structPreambleTemplate.Execute(outf, &str); err != nil {
			return err
		}
		if str.NoTable {
			continue
		}
		if err := structTemplate.Execute(outf, &str); err != nil {
			return err
		}
	}

	return nil
}
