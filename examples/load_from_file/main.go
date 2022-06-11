package main

import (
	"fmt"

	"github.com/S4eedb/dasl"
)

func main() {

	out, _ := dasl.GetSourcesFromFile("/home/babaee/workspace/dasl/testdata/test_source.list")
	fmt.Print(len(out))
}
