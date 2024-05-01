package pdfToImage

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

type pdfService struct{}

func NewPDFService() *pdfService {
	return &pdfService{}
}

func (s *pdfService) ConvertPDFToImages(pdfData []byte) ([]string, error) {
	// Create a temporary directory to store the images
	tempDir, err := os.MkdirTemp("", "pdf_images_")
	if err != nil {
		return nil, fmt.Errorf("failed to create temporary directory: %s", err.Error())
	}
	defer os.RemoveAll(tempDir) // Clean up temporary directory when done

	// Write PDF data to a temporary file
	pdfPath := filepath.Join(tempDir, "input.pdf")
	err = os.WriteFile(pdfPath, pdfData, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to write PDF data to file: %s", err.Error())
	}

	// Use pdfcpu to convert PDF to images
	imagePaths, err := convertPDFToImagesWithPDFCPU(pdfPath, tempDir)
	if err != nil {
		return nil, fmt.Errorf("failed to convert PDF to images: %s", err.Error())
	}

	return imagePaths, nil
}

// convertPDFToImagesWithPDFCPU converts PDF to images using pdfcpu package
func convertPDFToImagesWithPDFCPU(pdfPath, outputDir string) ([]string, error) {
	// Perform PDF to image conversion
	err := api.ExtractImagesFile(pdfPath, outputDir, nil, nil)
	if err != nil {
		return nil, err
	}

	// Get list of image files in the output directory
	files, err := os.ReadDir(outputDir)
	if err != nil {
		return nil, err
	}

	// Collect paths of the image files
	var imagePaths []string
	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".jpg") {
			imagePaths = append(imagePaths, filepath.Join(outputDir, file.Name()))
		}
	}

	return imagePaths, nil
}
