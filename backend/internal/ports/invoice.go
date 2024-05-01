package ports

import "github.com/gofiber/fiber/v2"

type InvoiceDBPort interface {
	CreateInvoice()
}

type InvoiceHandlerPort interface {
	CreateInvoice(c *fiber.Ctx) error
}
