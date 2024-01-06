package main

import (
	"emreddit/migration_manager"
	"flag"
)

var fileName string

func main() {
	flag.StringVar(&fileName, "name", "", "")
	flag.Parse()
	migration_manager.Init(fileName)
}
