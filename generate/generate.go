package generate

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func guessColumnType(goType string) (string, error) {
	if goType == "int" {
		return "qb.Int()", nil
	} else if goType == "int64" {
		return "qb.BigInt()", nil
	} else if goType == "string" {
		return "qb.Varchar()", nil
	} else if goType == "*string" {
		return "qb.Varchar()", nil
	} else if goType == "bool" {
		return "qb.Boolean()", nil
	} else if goType == "time.Time" {
		return "qb.Timestamp()", nil
	} else if goType == "*time.Time" {
		return "qb.Timestamp()", nil
	} else if goType == "uuid.UUID" {
		return "qb.UUID()", nil
	} else if goType == "uuid.NullUUID" {
		return "qb.UUID()", nil
	}
	return "", fmt.Errorf("Cannot guess column type for go type %s", goType)
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
	} else if goType == "uuid.NullUUID" {
		return "(uuid.NullUUID{})"
	}
	panic(fmt.Sprintf("I have no empty value for type '%v'", goType))
}

func prepareFieldData(str *StructData, f *FieldData) {
	var err error
	if f.ColumnName == "" {
		f.ColumnName = makeColumnName(f.Name)
	}
	if f.ColumnType == "" {
		f.ColumnType, err = guessColumnType(f.Type)
		if err != nil {
			panic(fmt.Sprintf("Failure on field '%s': Got err '%s'",
				f.Name, err))
		}
	}
	if f.ColumnModifiers == "" {
		if f.Tags.PrimaryKey {
			f.ColumnModifiers += ".PrimaryKey()"
		}
		if f.Tags.AutoIncrement {
			f.ColumnModifiers += ".AutoIncrement()"
			str.AutoIncrementPKey = f
		}
		if f.Tags.Null {
			f.ColumnModifiers += ".Null()"
		} else if f.Tags.NotNull {
			f.ColumnModifiers += ".NotNull()"
		} else if f.Type[0] == '*' || strings.Contains(f.Type, "Null") {
			f.ColumnModifiers += ".Null()"
		} else {
			f.ColumnModifiers += ".NotNull()"
		}
	}
	if f.EmptyValue == "" && f.Tags.PrimaryKey {
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

func loadEmbedded(path string, structs map[string]*StructData, fd FileData) []*StructData {
	var newStructs []*StructData

	var hasEmbed bool

	for _, str := range structs {
		if len(str.Embed) != 0 {
			hasEmbed = true
		}
	}

	if !hasEmbed {
		return []*StructData{}
	}

	otherStructsByName := make(map[string]*StructData)

	var err error
	newStructs, err = ParseDir(path)
	if err != nil {
		panic(err)
	}
	for _, str := range newStructs {
		prepareStructData(str, fd)
		otherStructsByName[str.Name] = str
	}

	allStructs := []*StructData{}
	for _, str := range structs {
		allStructs = append(allStructs, str)
	}
	allStructs = append(allStructs, newStructs...)

	for _, str := range allStructs {
		for _, name := range str.Embed {
			embedded, ok := structs[name]
			if !ok {
				embedded, ok = otherStructsByName[name]
			}
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
			} else {
				fmt.Println(
					"Could not find embedded struct definition for '" + name + "'")
			}
		}
	}
	return newStructs
}

func postPrepare(filedata *FileData, structs map[string]*StructData) {
	for _, str := range structs {
		for i := range str.Fields {
			if str.Fields[i].Tags.PrimaryKey {
				str.PKeyFields = append(str.PKeyFields, &str.Fields[i])
			}
		}
		if len(str.PKeyFields) == 0 {
			panic(fmt.Sprintf("No Primary Key found on %s", str.Name))
		}
	}
	for _, str := range structs {
		if str.Imported {
			continue
		}
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
func ProcessFile(logger *log.Logger, path string, file string, pack string, output string, fmt bool) error {

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
		if !str.NoTable {
			filedata.HasTables = true
		}
	}
	otherStructs := loadEmbedded(path, structsByName, filedata)
	for _, str := range otherStructs {
		if _, ok := structsByName[str.Name]; !ok {
			str.Imported = true
			structsByName[str.Name] = str
		}
	}
	postPrepare(&filedata, structsByName)

	outf, err := os.Create(output)
	if err != nil {
		return err
	}

	{
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
	}

	if fmt {
		cmd := exec.Command("gofmt", "-s", "-w", output)
		if err := cmd.Run(); err != nil {
			logger.Fatal(err)
		}
	}

	return nil
}
