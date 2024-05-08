# go-mdparser

go-mdparser is a tool that provides a set of commands that allow you to parse markdown files and generate html files.

## Usage

Run the following command in your terminal in the root of the project directory:

```
make parse
```
This command will run the `go-mdparser` tool with the default directories for the markdown files and the generated HTML files.

### Command Line Flags

The `go-mdparser` tool accepts the following command line flags:

- `md`: This flag allows you to specify the directory containing the Markdown files that you want to parse. The default value is `./markdown`.

- `out`: This flag allows you to specify the directory where the generated HTML files will be stored. The default value is `./generated`.

For example, if you want to parse the Markdown files in a directory named `my_markdown_files` and store the generated HTML files in a directory named `my_html_files`, you can run the following command:

```
make parse --md my_markdown_files --out my_html_files
```

