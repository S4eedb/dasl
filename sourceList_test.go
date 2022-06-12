package dasl_test

import (
	"testing"

	"github.com/S4eedb/dasl"
)

func TestGetSourcesFromFileNotExist(t *testing.T) {
	_, err := dasl.GetSourcesFromFile("file_not_exist.list")

	if err != dasl.ErrFilesnotexist {
		t.Fatalf(`%v, want "", error`, err)
	}
}

func TestGetSourcesFromFile(t *testing.T) {
	out, err := dasl.GetSourcesFromFile("/home/babaee/workspace/dasl/testdata/test_source.list")

	want := dasl.Source{
		Component:     "deb",
		ArchiveType:   "buster",
		RepositoryURL: "http://deb.debian.org/debian",
		Distributions: []string{"main", "contrib", "non-free"},
	}
	if err == nil {
		if want.Component != out[0].Component {
			t.Error("not equal sources list")
		}
		if want.ArchiveType != out[0].ArchiveType {
			t.Error("not equal ArchiveType")
		}
		if want.RepositoryURL != out[0].RepositoryURL {
			t.Error("not equal RepositoryURL")
		}
	}
}

func TestGetSourcesFromEmptyFile(t *testing.T) {
	_, err := dasl.GetSourcesFromFile("testdata/emptyfile.list")
	if err != dasl.ErrNoSourcesFound {
		t.Errorf("error is not %s", dasl.ErrNoSourcesFound)

	}
}

func TestParseSourceLine(t *testing.T) {
	dasl.ParseSourceLine("deb http://deb.debian.org/debian buster-updates main contrib non-free")
}
