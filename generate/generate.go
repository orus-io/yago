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
	} else {
		panic(fmt.Sprintf("I have no empty value for type '%v'", goType))
	}
}

func prepareFieldData(f *FieldData) {
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
		}
	}
	if f.EmptyValue == "" {
		f.EmptyValue = getEmptyValue(f.Type)
	}
}

func prepareStructData(str *StructData, fd FileData) {
	str.File = fd
	str.PrivateBasename = strings.ToLower(str.Name[0:1]) + str.Name[1:]
	for i := range str.Fields {
		prepareFieldData(&str.Fields[i])
		if str.Fields[i].Tags.PrimaryKey {
			str.PKeyFields = append(str.PKeyFields, &str.Fields[i])
		}
	}
}

// ProcessFile processes a go file and generates mapper and mappedstruct
// interfaces implementations for the yago structs.
func ProcessFile(logger *log.Logger, path string, file string, pack string) error {

	ext := filepath.Ext(file)
	base := strings.TrimSuffix(file, ext)

	outf, err := os.Create(filepath.Join(path, base+"_yago"+ext))
	if err != nil {
		return err
	}

	defer outf.Close()

	fd := FileData{Package: pack}
	if err := prologTemplate.Execute(outf, &fd); err != nil {
		return err
	}

	structs, err := ParseFile(filepath.Join(path, file))
	if err != nil {
		return err
	}
	for _, str := range structs {
		prepareStructData(str, fd)
		if err := structTemplate.Execute(outf, &str); err != nil {
			return err
		}
	}

	return nil
}
