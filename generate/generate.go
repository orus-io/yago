package generate

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
		if err := structTemplate.Execute(outf, &str); err != nil {
			return err
		}
	}

	return nil
}
