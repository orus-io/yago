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
}

func prepareStructData(str *StructData) {
	for i := range str.Fields {
		prepareFieldData(&str.Fields[i])
	}
}

// ProcessFile processes a go file and generates mapper and mappedstruct
// interfaces implementations for the yagorm structs.
func ProcessFile(logger *log.Logger, path string, file string, pack string) error {

	ext := filepath.Ext(file)
	base := strings.TrimSuffix(file, ext)

	outf, err := os.Create(filepath.Join(path, base+"_yagorm"+ext))
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
		prepareStructData(str)
		if err := structTemplate.Execute(outf, &str); err != nil {
			return err
		}
	}

	return nil
}
