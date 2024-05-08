package parser

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

func parseMarkdown(input string) string {
	scanner := bufio.NewScanner(strings.NewReader(input))
	var output strings.Builder
	var currentParagraph strings.Builder
	hashParagraphRegex := regexp.MustCompile(`^(\#{7,})\s*(.*)$`)
	regularParagraphRegex := regexp.MustCompile(`^([^#].*)`)

	for scanner.Scan() {
		line := scanner.Text()
		if hashParagraphRegex.Match([]byte(line)) || regularParagraphRegex.Match([]byte(line)) {
			if currentParagraph.Len() > 0 {
				currentParagraph.WriteString("\n")
			}
			currentParagraph.WriteString(line)
		} else if line == "" {
			if currentParagraph.Len() > 0 {
				output.WriteString(buildHtmlParagraph(currentParagraph.String()) + "\n")
				currentParagraph.Reset()
			}
			output.WriteString("\n")
		} else {
			if currentParagraph.Len() > 0 {
				output.WriteString(buildHtmlParagraph(currentParagraph.String()) + "\n")
				currentParagraph.Reset()
			}
			output.WriteString(buildHtmlHeading(line) + "\n")
		}
	}

	if currentParagraph.Len() > 0 {
		output.WriteString(buildHtmlParagraph(currentParagraph.String()) + "\n")
	}

	return output.String()
}

func buildHtmlHeading(input string) string {
	input = buildHtmlLink(input)
	headingRegex := regexp.MustCompile(`^(\#{1,6})\s*(.*)$`)
	matches := headingRegex.FindStringSubmatch(input)
	if len(matches) < 3 {
		return buildHtmlParagraph(input)
	}
	headerLevel := len(matches[1])
	headerText := matches[2]
	return fmt.Sprintf("<h%d>%s</h%d>", headerLevel, headerText, headerLevel)
}

func buildHtmlParagraph(input string) string {
	input = buildHtmlLink(input)
	return fmt.Sprintf("<p>%s</p>", input)
}

func buildHtmlLink(input string) string {
	linkRegex := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)

	input = linkRegex.ReplaceAllStringFunc(input, func(match string) string {
		matches := linkRegex.FindStringSubmatch(match)
		if len(matches) >= 3 {
			text := matches[1]
			url := matches[2]
			return fmt.Sprintf(`<a href="%s">%s</a>`, url, text)
		}
		return match
	})
	return input
}
