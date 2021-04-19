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
	inputName := flag.String("i", "", "file to convert")
	outputName := flag.String("o", "", "outpufilename")
	showVersion := flag.Bool("v", false, "version")
	flag.Parse()

	if *showVersion {
		fmt.Println(version)
		return
	}

	err := convert(*inputName, *outputName)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func convert(inputName, outputName string) (err error) {
	if inputName == "" {
		return errors.New("missing input file")
	}

	if outputName == "" {
		nameOnly := strings.TrimSuffix(inputName, filepath.Ext(inputName))
		outputName = nameOnly + ".pdf"
	}

	outputFile, err := os.Create(outputName)
	if err != nil {
		return fmt.Errorf("can't create outputfile %w", err)
	}
	defer outputFile.Close()

	reader, err := zip.OpenReader(inputName)
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
