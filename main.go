package main

import (
	"archive/zip"
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	rm2pdf "github.com/poundifdef/go-remarkable2pdf"
)

var version string

func main() {
	err := innerMain()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func innerMain() (err error) {
	var outputName string
	inputName := flag.String("i", "", "file to convert")
	flag.StringVar(&outputName, "o", "", "outpufilename")
	showVersion := flag.Bool("v", false, "version")
	// outputFormat := flag.String("f","pdf", "format (pdf, png)")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	if *inputName == "" {
		return errors.New("missing input file")
	}

	if outputName == "" {
		nameOnly := strings.TrimSuffix(*inputName, filepath.Ext(*inputName))
		outputName = nameOnly + ".pdf"
	}

	outputFile, err := os.Create(outputName)
	if err != nil {
		return fmt.Errorf("can't create outputfile %w", err)
	}
	defer outputFile.Close()

	reader, err := zip.OpenReader(*inputName)
	if err != nil {
		return fmt.Errorf("can't open file %w", err)
	}
	defer reader.Close()

	err = rm2pdf.RenderRmNotebookFromZip(&reader.Reader, outputFile)
	if err != nil {
		return fmt.Errorf("can't open file %w", err)
	}
	return nil

}
