package main

import (
	"emreddit/migrate"
	"flag"
)

var fileName string

func main() {
	flag.StringVar(&fileName, "name", "", "")
	flag.Parse()
	migrate.Init(fileName)
}
