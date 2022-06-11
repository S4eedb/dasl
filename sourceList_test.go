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
	out, _ := dasl.GetSourcesFromFile("testdata/test_source_one_line.list")
	want := dasl.Source{
		Component: "deb",
	}
	if want.Component != out[0].Component {
		t.Error("")
	}
}

func TestParseSourceLine(t *testing.T) {
	dasl.ParseSourceLine("deb http://deb.debian.org/debian buster-updates main contrib non-free")
}
