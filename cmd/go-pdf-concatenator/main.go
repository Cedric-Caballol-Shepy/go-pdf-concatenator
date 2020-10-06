package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	pdfcpu "github.com/pdfcpu/pdfcpu/pkg/api"
)

func main() {
	var filesToConcatenate []string

	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		if filepath.Ext(path) == ".pdf" {
			filesToConcatenate = append(filesToConcatenate, path)
		}
		return nil
	})
	if err == nil && len(filesToConcatenate) > 0 {
		var outputName string
		fmt.Println("Choose the pdf file name...")
		_, err = fmt.Scan(&outputName)
		if err == nil {
			err = pdfcpu.MergeCreateFile(filesToConcatenate, outputName+".pdf", nil)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
}
