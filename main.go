package main

import (
	"flag"
	"fmt"
	rm2pdf "github.com/poundifdef/go-remarkable2pdf"
	"os"
	"path/filepath"
	"strings"
)

var version string

func main() {
	var outputName string
	inputName := flag.String("i", "", "file to convert")
	flag.StringVar(&outputName, "o", "", "outpufilename")
	showVersion := flag.Bool("v", false, "version")
	// outputFormat := flag.String("f","pdf", "format (pdf, png)")
	flag.Parse()
	if *showVersion {
		fmt.Println(version)
		os.Exit(0)
	}

	if *inputName == "" {
		fmt.Fprintln(os.Stderr, "missing input file")
		os.Exit(1)
	}

	if outputName == "" {
		nameOnly := strings.TrimSuffix(*inputName, filepath.Ext(*inputName))
		outputName = nameOnly + ".pdf"
	}

	outputFile, err := os.Create(outputName)
	if err != nil {
		fmt.Fprintln(os.Stderr, "can't create outputfile ", err)

		os.Exit(1)
	}

	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintln(os.Stderr, "can't convert ", r)
			os.Exit(1)
		}
	}()
	rm2pdf.RenderRmNotebook(*inputName, outputFile)
	outputFile.Close()

}
