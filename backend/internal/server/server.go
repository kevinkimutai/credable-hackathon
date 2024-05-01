package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/kevinkimutai/credablehackathon/internal/ports"
)

type ServerAdapter struct {
	port    string
	invoice ports.InvoiceHandlerPort
}

func New(
	port string,
	invoice ports.InvoiceHandlerPort) *ServerAdapter {
	return &ServerAdapter{
		port:    port,
		invoice: invoice}
}

func (s *ServerAdapter) Run() {
	//Initialize Fiber
	app := fiber.New()

	//Logger Middleware
	app.Use(logger.New())

	//Swagger Middleware
	// cfg := swagger.Config{
	// 	BasePath: "/api/v1/",
	// 	FilePath: "./docs/swagger/swagger.json",
	// 	Path:     "swagger",
	// 	Title:    "Swagger Movie API Docs",
	// }
	// app.Use(swagger.New(cfg))

	// Define routes
	//app.Route("/api/v1/customer", s.CustomerRouter)
	app.Route("/api/v1/invoice", s.InvoiceRouter)
	//app.Route("/api/v1/product", s.ProductRouter)

	app.Listen(":" + s.port)
}
