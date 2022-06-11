package main

import (
	"fmt"

	"github.com/S4eedb/dasl"
)

func main() {

	source, err := dasl.ParseSourceLine("deb http://deb.debian.org/debian buster-updates main contrib non-free")
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print(source)
	}
}
