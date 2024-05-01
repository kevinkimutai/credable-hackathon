package ports

type PDFPort interface {
	ConvertPDFToImages(pdfData []byte) ([]string, error)
}
