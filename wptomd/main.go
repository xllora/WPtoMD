// Binary wptomd reads a WordPress XML export file and converts into
// markdown files.
package main

import (
	"fmt"
	"os"

	"github.com/xllora/WPtoMD/convert"
	"github.com/xllora/WPtoMD/io"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Fprintf(os.Stderr, "Wrong number of arguments provided\n\n\t Usage: wptomd <source_xml_file> <destination_folder>")
		os.Exit(1)
	}
	xml, err := io.ReadFile(args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read XML export file %q with error: %v", args[1], err)
		os.Exit(1)
	}
	fmt.Fprintf(os.Stdout, "Read %d bytes from the WP XML export dump...\n", len(xml))
	data, err := convert.ToMarkDown(xml)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to convert XML export with error: %v", err)
		os.Exit(1)
	}
	fmt.Println(data)
}