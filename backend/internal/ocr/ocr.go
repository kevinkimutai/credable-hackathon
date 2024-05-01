// ocr/service.go
package ocr

import (
	"github.com/otiai10/gosseract/v2"
)

type Service interface {
	GetTextFromImage(imagePath string) (string, error)
}

type ocrService struct{}

func NewOCRService() Service {
	return &ocrService{}
}

func (s *ocrService) GetTextFromImage(imagePath string) (string, error) {
	// Create a new Tesseract client
	client := gosseract.NewClient()
	defer client.Close()

	// Set the image file
	client.SetImage(imagePath)

	// Perform OCR
	text, err := client.Text()
	if err != nil {
		return "", err
	}

	return text, nil
}
