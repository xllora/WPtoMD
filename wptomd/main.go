// Binary wptomd reads a WordPress XML export file and converts into
// markdown files.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"

	"github.com/xllora/WPtoMD/convert"
	"github.com/xllora/WPtoMD/io"
	"github.com/xllora/WPtoMD/markdown"
)

var (
	frontMatterFormat = flag.String("front_matter_format", "TOML", "Format of the front matter {JSON, TOML, YAML}")
)

func main() {
	flag.Parse()
	args, fmf := os.Args, markdown.TOML
	switch *frontMatterFormat {
	case "JSON":
		fmf = markdown.JSON
	case "TOML":
		fmf = markdown.TOML
	case "YAML":
		fmf = markdown.YAML
	default:
		fmt.Fprintf(os.Stderr, "Wrong front matter format %q; valid formats JSON, TOML, YAML", *frontMatterFormat)
		os.Exit(1)
	}
	if len(args) < 3 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments provided\n\n\t Usage: wptomd [--front_matter_format={JSON, TOML, YAML}] <source_xml_file> <destination_folder>")
		os.Exit(1)
	}
	xml, err := io.ReadFile(args[len(args)-2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read XML export file %q with error: %v", args[len(args)-2], err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Read %d bytes from the WP XML export dump...\n", len(xml))
	data, err := convert.ToMarkDown(xml)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert XML export with error: %v", err)
		os.Exit(1)
	}
	buff := bytes.NewBufferString("")
	for _, item := range data.Channel.Items {
		if err := markdown.ToFrontMatter(buff, fmf, &item); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to convert item to MD with error: %v", err)
			os.Exit(1)
		}
	}
	fmt.Println(buff.String())
}
