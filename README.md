## Stack

- **Backend:** Golang
- **Frontend:** Next.js

## Approach

1. **Upload PDF to Firebase:** 
   - Users upload PDF files to Firebase Storage.
   
2. **Get URL and Send to Backend:**
   - Retrieve the URL from firebase of the uploaded PDF file.
   - Send the URL to the backend.

3. **Golang Backend:**
   - Receives the PDF URL from the frontend.
   - Downloads the PDF file from the provided URL.
   - Extracts text from the PDF using Tesseract OCR (Optical Character Recognition).
   - Validates the extracted texts.
   - Retrieves key texts from the extracted data.
   - Returns the pro
