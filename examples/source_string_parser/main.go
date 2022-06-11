package main

import (
	"github.com/S4eedb/dasl"
)

func main() {

	dasl.ParseSourceLine("deb http://deb.debian.org/debian buster-updates main contrib non-free")
}
