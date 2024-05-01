// storage/client.go
package storage

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type StorageClient struct{}

func NewStorageClient() *StorageClient {
	return &StorageClient{}
}

func (c *StorageClient) DownloadPDF(pdfURL string) ([]byte, error) {
	// Send HTTP GET request to download the PDF
	resp, err := http.Get(pdfURL)
	if err != nil {
		return nil, fmt.Errorf("failed to download PDF: %s", err.Error())
	}
	defer resp.Body.Close()

	// Check response status code
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download PDF: unexpected status code: %d", resp.StatusCode)
	}

	// Read PDF content from response body
	pdfData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read PDF data: %s", err.Error())
	}

	return pdfData, nil
}
