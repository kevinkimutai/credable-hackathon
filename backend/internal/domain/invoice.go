package domain

type Invoice struct {
	InvoiceURL string `json:"invoice_url"`
}

type ErrorResponse struct {
	StatusCode uint   `json:"status_code"`
	Message    string `json:"message"`
}
