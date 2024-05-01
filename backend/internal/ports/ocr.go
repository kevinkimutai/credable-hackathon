package ports

type OCRPort interface {
	GetTextFromImage(imagePath string) (string, error)
}
