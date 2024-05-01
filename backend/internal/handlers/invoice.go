package handler

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kevinkimutai/credablehackathon/internal/domain"
	"github.com/kevinkimutai/credablehackathon/internal/ports"
)

type InvoiceService struct {
	//db      ports.InvoiceDBPort
	pdf     ports.PDFPort
	ocr     ports.OCRPort
	storage ports.StoragePort
}

func NewInvoiceService( //db ports.InvoiceDBPort,
	pdf ports.PDFPort,
	ocr ports.OCRPort,
	storage ports.StoragePort) *InvoiceService {
	return &InvoiceService{pdf: pdf, ocr: ocr, storage: storage}
}

func (s *InvoiceService) CreateInvoice(c *fiber.Ctx) error {
	newInvoice := domain.Invoice{}

	//Bind To struct
	if err := c.BodyParser(&newInvoice); err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    err.Error(),
			})
	}

	//Check missing values
	if newInvoice.InvoiceURL == "" {
		return c.Status(400).JSON(
			domain.ErrorResponse{
				StatusCode: 400,
				Message:    "Missing invoice_url field",
			})
	}

	// Download PDF
	pdfData, err := s.storage.DownloadPDF(newInvoice.InvoiceURL)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("Failed to download PDF: %s", err.Error()),
			})
	}

	// Convert PDF to images
	imagePaths, err := s.pdf.ConvertPDFToImages(pdfData)
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("Failed to convert PDF to images: %s", err.Error()),
			})
	}

	///OCR GET TEXT
	// OCR: Get text from each image
	var extractedTexts []string
	for _, imagePath := range imagePaths {
		text, err := s.ocr.GetTextFromImage(imagePath)
		if err != nil {
			return c.Status(500).JSON(
				domain.ErrorResponse{
					StatusCode: 500,
					Message:    fmt.Sprintf("Failed to perform OCR on image: %s", err.Error()),
				})
		}
		extractedTexts = append(extractedTexts, text)
	}

	fmt.Println("EXTRACTED TEXTS", extractedTexts)
	// Validate extracted text
	for _, text := range extractedTexts {
		valid, err := ValidateInvoice(text)
		if err != nil {
			return c.Status(500).JSON(
				domain.ErrorResponse{
					StatusCode: 500,
					Message:    fmt.Sprintf("Validation error: %s", err.Error()),
				})
		}
		if !valid {
			return c.Status(400).JSON(
				domain.ErrorResponse{
					StatusCode: 400,
					Message:    "Invalid invoice",
				})
		}
	}

	// Parse key information
	var schoolName, invoiceNumber, invoiceAmount, parentsName, schoolAccountNumber string
	schoolName, invoiceNumber, invoiceAmount, parentsName, schoolAccountNumber, err = ParseInvoice(strings.Join(extractedTexts, " "))
	if err != nil {
		return c.Status(500).JSON(
			domain.ErrorResponse{
				StatusCode: 500,
				Message:    fmt.Sprintf("Error parsing invoice: %s", err.Error()),
			})
	}

	// Return key information
	return c.Status(200).JSON(fiber.Map{
		"school_name":           schoolName,
		"invoice_number":        invoiceNumber,
		"invoice_amount":        invoiceAmount,
		"parents_name":          parentsName,
		"school_account_number": schoolAccountNumber,
	})

}

func ValidateInvoice(extractedText string) (bool, error) {
	// Check if the extracted text contains specific keywords or patterns
	if !containsKeyword(extractedText, "Invoice Number") {
		return false, fmt.Errorf("Invoice number not found")
	}

	// Validate date format (assuming date is in YYYY-MM-DD format)
	if !isValidDateFormat(extractedText, "YYYY-MM-DD") {
		return false, fmt.Errorf("Invalid date format")
	}

	// Validate total amount format (assuming it's a numeric value)
	if !isValidAmountFormat(extractedText) {
		return false, fmt.Errorf("Invalid total amount format")
	}

	// Additional validation criteria...

	// If all validations pass, return true
	return true, nil
}

// containsKeyword checks if the extracted text contains a specific keyword or pattern
func containsKeyword(text, keyword string) bool {
	// Implement logic to check if the keyword exists in the text
	return strings.Contains(text, keyword)
}

// isValidDateFormat checks if the extracted text contains a date in the specified format
func isValidDateFormat(text, format string) bool {
	// Implement logic to validate the date format using regular expressions or other techniques
	// For simplicity, we'll assume any string of the form "YYYY-MM-DD" is a valid date format
	datePattern := `\d{4}-\d{2}-\d{2}`
	match, _ := regexp.MatchString(datePattern, text)
	return match
}

// isValidAmountFormat checks if the extracted text contains a valid numeric amount
func isValidAmountFormat(text string) bool {
	// Implement logic to validate the amount format using regular expressions or other techniques
	// For simplicity, we'll assume any string containing digits and a currency symbol is a valid amount format
	amountPattern := `\d+\.\d+ \$`
	match, _ := regexp.MatchString(amountPattern, text)
	return match
}

// ParseInvoice extracts key information from the extracted text
func ParseInvoice(extractedText string) (schoolName, invoiceNumber, invoiceAmount, parentsName, schoolAccountNumber string, err error) {
	// Implement parsing logic to extract key information from the extracted text
	schoolName, err = extractSchoolName(extractedText)
	if err != nil {
		return "", "", "", "", "", err
	}

	invoiceNumber, err = extractInvoiceNumber(extractedText)
	if err != nil {
		return "", "", "", "", "", err
	}

	invoiceAmount, err = extractInvoiceAmount(extractedText)
	if err != nil {
		return "", "", "", "", "", err
	}

	parentsName, err = extractParentsName(extractedText)
	if err != nil {
		return "", "", "", "", "", err
	}

	schoolAccountNumber, err = extractSchoolAccountNumber(extractedText)
	if err != nil {
		return "", "", "", "", "", err
	}

	// Return extracted key information
	return schoolName, invoiceNumber, invoiceAmount, parentsName, schoolAccountNumber, nil
}

// Implement functions to extract key information from the extracted text
// These functions will use regular expressions or other techniques to search for specific patterns or keywords

func extractSchoolName(text string) (string, error) {
	// Implement logic to extract school name from the text
	// Example: Search for a pattern or keyword indicating the school name
	return "", nil
}

func extractInvoiceNumber(text string) (string, error) {
	// Implement logic to extract invoice number from the text
	// Example: Search for a pattern or keyword indicating the invoice number
	return "", nil
}

func extractInvoiceAmount(text string) (string, error) {
	// Implement logic to extract invoice amount from the text
	// Example: Search for a pattern or keyword indicating the invoice amount
	return "", nil
}

func extractParentsName(text string) (string, error) {
	// Implement logic to extract parents' name from the text
	// Example: Search for a pattern or keyword indicating the parents' name
	return "", nil
}

func extractSchoolAccountNumber(text string) (string, error) {
	// Implement logic to extract school account number from the text
	// Example: Search for a pattern or keyword indicating the school account number
	return "", nil
}

// func (s *ProductService) GetAllProducts(c *fiber.Ctx) error {
// 	//Get Query Params
// 	m := c.Queries()

// 	//Bind To ProductParams
// 	prodParams := domain.CheckProductsParams(m)

// 	//Get All Products API
// 	prod, err := s.api.GetAllProducts(prodParams)
// 	if err != nil {
// 		return c.Status(500).JSON(
// 			domain.ErrorResponse{
// 				StatusCode: 500,
// 				Message:    err.Error(),
// 			})

// 	}
// 	return c.Status(200).JSON(
// 		domain.ProductsResponse{
// 			StatusCode:    200,
// 			Message:       "Successfully retrieved products",
// 			Page:          prod.Page,
// 			NumberOfPages: prod.NumberOfPages,
// 			Total:         prod.Total,
// 			Data:          prod.Data,
// 		})
// }

// func (s *ProductService) GetProductByID(c *fiber.Ctx) error {
// 	productID := c.Params("productID")

// 	//Get Product API
// 	product, err := s.api.GetProduct(productID)
// 	if err != nil {
// 		return c.Status(500).JSON(
// 			domain.ErrorResponse{
// 				StatusCode: 500,
// 				Message:    err.Error(),
// 			})
// 	}

// 	return c.Status(200).JSON(
// 		domain.ProductResponse{
// 			StatusCode: 200,
// 			Message:    "Successfully retrieved product",
// 			Data:       product})
// }

// func (s *ProductService) DeleteProduct(c *fiber.Ctx) error {
// 	productID := c.Params("productID")

// 	//Delete Product API
// 	err := s.api.DeleteProduct(productID)
// 	if err != nil {
// 		return c.Status(500).JSON(
// 			domain.ErrorResponse{
// 				StatusCode: 500,
// 				Message:    err.Error(),
// 			})
// 	}

// 	return c.Status(204).JSON(
// 		domain.CustomerResponse{
// 			StatusCode: 204,
// 			Message:    "Successfully Deleted product",
// 		})

// }
