package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	handler "github.com/kevinkimutai/credablehackathon/internal/handlers"
	"github.com/kevinkimutai/credablehackathon/internal/ocr"
	"github.com/kevinkimutai/credablehackathon/internal/pdfToImage"
	"github.com/kevinkimutai/credablehackathon/internal/server"
	"github.com/kevinkimutai/credablehackathon/internal/storage"
)

func main() {
	// Init Environment Variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env files")
	}

	// //Env Variables
	APP_PORT := os.Getenv("APP_PORT")
	// POSTGRES_USERNAME := os.Getenv("POSTGRES_USERNAME")
	// POSTGRES_PASSWORD := os.Getenv("POSTGRES_PASSWORD")
	// DATABASE_PORT := os.Getenv("DB_PORT")

	// DBURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
	// 	POSTGRES_USERNAME,
	// 	POSTGRES_PASSWORD,
	// 	"localhost",
	// 	DATABASE_PORT,
	// 	"invoicedb")

	// Firebase Storage configuration
	//bucketName := "your-firebase-storage-bucket"
	//localPDFPath := "local/path/to/save/pdf.pdf"
	//outputPath := "output/path/to/save/images"

	//Dependency injection
	//Connect To DB
	//dbAdapter := db.NewDB(DBURL)

	//OCR
	ocrService := ocr.NewOCRService()

	//PDFTOImage
	pdf := pdfToImage.NewPDFService()

	//Storage
	storage := storage.NewStorageClient()

	//userService := handler.NewUserService(customerRepo)
	invoiceService := handler.NewInvoiceService(
		//dbAdapter,
		pdf,
		ocrService,
		storage)
	//schoolsService:=handler.NewSchools

	// authService, err := auth.New(dbAdapter)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	//Server
	server := server.New(
		APP_PORT,
		invoiceService,
	)
	server.Run()

}
