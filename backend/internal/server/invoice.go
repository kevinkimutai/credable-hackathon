package server

import "github.com/gofiber/fiber/v2"

func (s *ServerAdapter) InvoiceRouter(api fiber.Router) {
	api.Post("/", s.invoice.CreateInvoice)
	// api.Get("/", s.auth.IsAuthenticated, s.product.GetAllProducts)
	// api.Get("/:productID", s.auth.IsAuthenticated, s.product.GetProductByID)
	// api.Delete("/:productID", s.auth.IsAuthenticated, s.auth.AllowedRoles("Admin"), s.product.DeleteProduct)

}
