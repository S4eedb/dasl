package dasl

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
)

var ErrLineIsEmpty error = errors.New("line is Empty")

type Source struct {
	Component     string
	ArchiveType   string
	RepositoryURL string
	Distributions []string
}

type SourcesList struct {
	Sources []Source
}

var _ io.Reader = (*os.File)(nil)

func NewSources(SourceListString string) (SourcesList, error) {
	var sourcesList SourcesList
	scanner := bufio.NewScanner(strings.NewReader(SourceListString))
	scanner.Scan()
	lines := readBody(scanner)
	for _, line := range lines {
		ParsedSource, _ := ParseSourceLine(line)
		sourcesList.Sources = append(sourcesList.Sources, ParsedSource)

	}
	return sourcesList, nil
}

func readBody(scanner *bufio.Scanner) []string {
	return strings.Split(scanner.Text(), "\n")
}

func ParseSourceLine(SourceString string) (Source, error) {
	words := strings.Fields(SourceString)
	if len(words) == 0 {
		return Source{}, ErrLineIsEmpty
	}
	var newWord []string
	for _, word := range words {
		fistChar := strings.Index(word, "[")
		if fistChar != 0 {
			newWord = append(newWord, word)
		}
	}
	return Source{
		Component:     newWord[0],
		ArchiveType:   newWord[1],
		RepositoryURL: newWord[2],
		Distributions: newWord[3:],
	}, nil
}
