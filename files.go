package dasl

import (
	"errors"

	"github.com/spf13/afero"
)

var AppFs = afero.NewOsFs()
var ErrFilesnotexist error = errors.New("file not exist")

func GetSourcesFromFile(filePath string) ([]Source, error) {
	fs := afero.NewOsFs()
	fi, err := fs.Stat(filePath)
	isExist, _ := afero.Exists(fs, filePath)
	if !isExist {
		return nil, ErrFilesnotexist
	}

	if err != nil {
		return nil, err
	}
	fi.Size()
	buferFile, err := afero.ReadFile(fs, filePath)

	if err != nil {
		return nil, err
	}
	fileString := string(buferFile)
	return NewSources(fileString)
}
