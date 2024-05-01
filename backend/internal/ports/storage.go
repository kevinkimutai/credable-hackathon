package ports

type StoragePort interface {
	DownloadPDF(pdfURL string) ([]byte, error)
}
