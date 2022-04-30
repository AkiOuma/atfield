package main

import (
	"flag"
	"log"

	"github.com/AkiOuma/atfield/internal/atfield"
)

func main() {
	in := flag.String("in", ".", "path of link definition file")
	out := flag.String("out", ".", "path of exporting convert definition file")
	flag.Parse()
	if err := atfield.Unfold(*in, *out); err != nil {
		log.Panicf("atfield unfold failed because of: %v", err)
	}
}
