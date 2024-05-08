package main

import (
	"flag"

	"github.com/dlanell/go-mdparser/pkg/parser"
)

func main() {
	markdownDir := flag.String("md", "./markdown", "Directory containing Markdown files")
	outputDir := flag.String("out", "./generated", "Directory where HTML files will be stored")
	flag.Parse()

	parser.NewApp(*markdownDir, *outputDir).Run()
}
