package main

import (
	migrationutils "emreddit/migration-utils"
	"flag"
)

var FileName string

func main() {
	flag.StringVar(&FileName, "name", "", "")
	flag.Parse()
	migrationutils.Init(FileName)
}
