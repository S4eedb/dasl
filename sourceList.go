package dasl

import (
	"errors"
	"io"
	"os"
	"strings"
)

var ErrLineIsEmpty error = errors.New("line is Empty")
var ErrNoSourcesFound error = errors.New("no sources found")

type Source struct {
	Component     string
	ArchiveType   string
	RepositoryURL string
	Distributions []string
}

var _ io.Reader = (*os.File)(nil)

func NewSources(SourceListString string) ([]Source, error) {
	var sourcesList []Source
	lines := readBody(SourceListString)
	for _, line := range lines {
		ParsedSource, err := ParseSourceLine(line)
		if err != nil {
			continue
		}
		sourcesList = append(sourcesList, ParsedSource)

	}
	if len(sourcesList) == 0 {
		return nil, ErrNoSourcesFound
	}
	return sourcesList, nil
}

func readBody(SourceListString string) []string {
	lines := strings.Split(SourceListString, "\n")
	return lines
}

func ParseSourceLine(SourceString string) (Source, error) {
	words := strings.Fields(SourceString)
	if len(words) == 0 {
		return Source{}, ErrLineIsEmpty
	}
	var newWord []string
	for _, word := range words {
		if strings.Index(word, "#") == 0 {
			break
		}
		if strings.Index(word, "[") == 0 {
			continue
		}
		newWord = append(newWord, word)
	}
	if len(newWord) == 0 {
		return Source{}, ErrLineIsEmpty
	}
	return Source{
		Component:     newWord[0],
		ArchiveType:   newWord[1],
		RepositoryURL: newWord[2],
		Distributions: newWord[3:],
	}, nil
}
