package parser

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseMarkdown(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// TODO: Add test cases.
		{
			name: "given markdown with single line paragraph and heading, build html",
			input: `# Sample Document

Hello!

This is sample markdown for the [Mailchimp](https://www.mailchimp.com) homework assignment.`,
			want: `<h1>Sample Document</h1>

<p>Hello!</p>

<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment.</p>
`,
		},
		{
			name: "given markdown with multi line paragraph and multiple headings, build html",
			input: `# Header one

Hello there

How are you?
What's going on?

## Another Header

This is a paragraph [with an inline link](http://google.com). Neat, eh?

## This is a header [with a link](http://yahoo.com)`,
			want: `<h1>Header one</h1>

<p>Hello there</p>

<p>How are you?
What's going on?</p>

<h2>Another Header</h2>

<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>

<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, parseMarkdown(tt.input), "parseMarkdown(%v)", tt.input)
		})
	}
}

func Test_buildHtmlHeading(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// TODO: Add test cases.
		{
			"given markdown heading 1, return html heading",
			"# heading",
			`<h1>heading</h1>`,
		},
		{
			"given markdown heading 2, return html heading",
			"## heading",
			`<h2>heading</h2>`,
		},
		{
			"given markdown heading 3, return html heading",
			"### heading",
			`<h3>heading</h3>`,
		},
		{
			"given markdown heading 4, return html heading",
			"#### heading",
			`<h4>heading</h4>`,
		},
		{
			"given markdown heading 5, return html heading",
			"##### heading",
			`<h5>heading</h5>`,
		},
		{
			"given markdown heading 6, return html heading",
			"###### heading",
			`<h6>heading</h6>`,
		},
		{
			"given markdown heading 1 with link, return html heading with link",
			"# heading [Mailchimp](https://www.mailchimp.com)",
			`<h1>heading <a href="https://www.mailchimp.com">Mailchimp</a></h1>`,
		},
		{
			"given markdown heading 1 with 2 links, return html heading with links",
			"# heading [Mailchimp](https://www.mailchimp.com) [Intuit](https://www.intuit.com)",
			`<h1>heading <a href="https://www.mailchimp.com">Mailchimp</a> <a href="https://www.intuit.com">Intuit</a></h1>`,
		},
		{
			"given invalid heading return html paragraph",
			"heading",
			`<p>heading</p>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, buildHtmlHeading(tt.input), "buildHtmlHeading(%v)", tt.input)
		})
	}
}

func Test_buildHtmlParagraph(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		// TODO: Add test cases.
		{
			"given markdown paragraph 1 with link, return html paragraph with link",
			"paragraph [Mailchimp](https://www.mailchimp.com)",
			`<p>paragraph <a href="https://www.mailchimp.com">Mailchimp</a></p>`,
		},
		{
			"given markdown paragraph 1 with 2 links, return html paragraph with links",
			"paragraph [Mailchimp](https://www.mailchimp.com) [Intuit](https://www.intuit.com)",
			`<p>paragraph <a href="https://www.mailchimp.com">Mailchimp</a> <a href="https://www.intuit.com">Intuit</a></p>`,
		},
		{
			"given string starting with 7 hashes return html paragraph",
			"####### heading",
			`<p>####### heading</p>`,
		},
		{
			"given regular paragraph, html paragraph",
			"paragraph",
			`<p>paragraph</p>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, buildHtmlParagraph(tt.input), "buildHtmlParagraph(%v)", tt.input)
		})
	}
}
