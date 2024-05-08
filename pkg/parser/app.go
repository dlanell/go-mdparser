package parser

import (
	"os"
	"path/filepath"
	"strings"
)

type App struct {
	markdownDir string
	outputDir   string
}

func NewApp(markdownDir, outputDir string) *App {
	return &App{
		markdownDir: markdownDir,
		outputDir:   outputDir,
	}
}

func (a *App) Run() {
	files, err := filepath.Glob(filepath.Join(a.markdownDir, "*.md"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		markdownContent, err := readFile(file)
		if err != nil {
			panic(err)
		}
		htmlContent := parseMarkdown(markdownContent)
		baseFilename := strings.TrimSuffix(filepath.Base(file), filepath.Ext(file)) + ".html"
		outputFile := filepath.Join(a.outputDir, baseFilename)
		os.MkdirAll(a.outputDir, os.ModePerm)
		err = os.WriteFile(outputFile, []byte(htmlContent), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func readFile(file string) (string, error) {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(fileBytes), nil
}
