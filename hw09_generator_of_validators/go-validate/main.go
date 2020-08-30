package main

import (
	"go/format"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Filename was not found")
	}

	input, _ := parser.ParseFile(token.NewFileSet(), os.Args[1], nil, parser.AllErrors)
	buf, err := writeToTemplate(input)
	if err != nil {
		log.Fatal(err)
	}

	code, err := format.Source(buf.Bytes())
	if err != nil {
		log.Fatal("Cannot prettify code", err)
	}
	err = ioutil.WriteFile(strings.Replace(os.Args[1], ".go", "_validation_generated.go", 1), code, os.ModePerm)
	if err != nil {
		log.Fatal("Cannot write formatted file", err)
	}
}
