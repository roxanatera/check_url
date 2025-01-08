package main

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html/v2"
    "url_checker/handlers"
)



func main() {
   
	engine := html.New("templates", ".html") 


    app := fiber.New(fiber.Config{
        Views: engine, 
    })
	
    app.Get("/", func(c *fiber.Ctx) error {
        return c.Render("index", fiber.Map{
            "Title": "Comprobador de URLs",
        }) 
    })

    // Ruta para comprobar URLs
    app.Post("/check", func(c *fiber.Ctx) error {
        var request struct {
            URLs []string `json:"urls"`
        }

        // Parsear el cuerpo de la solicitud
        if err := c.BodyParser(&request); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
                "error": "Invalid request format",
            })
        }

        // Comprobar el estado de las URLs
        results := handlers.CheckURLs(request.URLs)
        return c.JSON(results)
    })

    // Iniciar el servidor
    app.Listen(":3000")
}
