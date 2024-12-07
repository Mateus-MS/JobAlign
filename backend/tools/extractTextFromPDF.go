package tools

import (
	"fmt"
	"os"

	"github.com/unidoc/unipdf/v3/extractor"
	"github.com/unidoc/unipdf/v3/model"
)

func ExtractTextFromPDF() string {
	f, err := os.Open("./testPDF.pdf")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	pdfReader, err := model.NewPdfReader(f)
	if err != nil {
		fmt.Println(err)
	}

	var text string
	page, err := pdfReader.GetPage(1)
	if err != nil {
		fmt.Println(err)
	}

	ex, err := extractor.New(page)
	if err != nil {
		fmt.Println(err)
	}

	text, _ = ex.ExtractText()
	return text
}
