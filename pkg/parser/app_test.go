package parser

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	type args struct {
		markdownDir string
		outputDir   string
	}
	tests := []struct {
		name string
		args args
		want *App
	}{
		// TODO: Add test cases.
		{
			name: "given markdown and output directory, return app",
			args: args{
				markdownDir: "markdown",
				outputDir:   "generated",
			},
			want: &App{
				markdownDir: "markdown",
				outputDir:   "generated",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, NewApp(tt.args.markdownDir, tt.args.outputDir), "NewApp(%v, %v)", tt.args.markdownDir, tt.args.outputDir)
		})
	}
}

func TestApp_Run(t *testing.T) {
	t.Run("should generate html files", func(t *testing.T) {
		markdownDir := "testdata"
		outputDir := "testdata/generated"

		app := NewApp(markdownDir, outputDir)
		app.Run()

		files, err := filepath.Glob(filepath.Join(outputDir, "*.html"))
		assert.Nil(t, err)
		assert.Equal(t, files, []string{"testdata/generated/test-1.html", "testdata/generated/test-2.html"})
	})
	t.Run("should generate valid html content in files", func(t *testing.T) {
		markdownDir := "testdata"
		outputDir := "testdata/generated"

		app := NewApp(markdownDir, outputDir)
		app.Run()

		expectedFiles := map[string]string{
			"testdata/generated/test-1.html": `<h1>Sample Document</h1>

<p>Hello!</p>

<p>This is sample markdown for the <a href="https://www.mailchimp.com">Mailchimp</a> homework assignment.</p>
`,
			"testdata/generated/test-2.html": `<h1>Header one</h1>

<p>Hello there</p>

<p>How are you?
What's going on?</p>

<h2>Another Header</h2>

<p>This is a paragraph <a href="http://google.com">with an inline link</a>. Neat, eh?</p>

<h2>This is a header <a href="http://yahoo.com">with a link</a></h2>
`,
		}

		for fileName, content := range expectedFiles {
			fileContent, err := readFile(fileName)
			assert.Nil(t, err)
			assert.Equal(t, content, fileContent)
		}

	})
}
