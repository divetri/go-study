package main

import "github.com/gofiber/fiber/v3"

//Returning JSON
type JSONTextResponse struct {
	Message string
}

func main() {
	var m = ""

	//Create New Instance
	app := fiber.New()

	//Basic
	// app.Get("/", func(c fiber.Ctx) error {
	// 	return c.SendString("Hello, World 👋!")
	// })

	//Parameter with Handling
	app.Get("/:account?", func(c fiber.Ctx) error {
		if c.Params("account") != "" {
			//return c.SendString("Hello, " + c.Params("account"))
			m = "Hello, " + c.Params("account"+"👋")
		} else {
			//return c.SendString("Hello?")
			m = "Hello?"
		}
		return c.JSON(JSONTextResponse{Message: m})
	})

	//Wildcards
	app.Get("/api/*", func(c fiber.Ctx) error {
		//return c.SendString("API path: " + c.Params("*"))
		m = "API path: " + c.Params("*")
		return c.JSON(JSONTextResponse{Message: m})
	})

	app.Listen(":3000")
}
