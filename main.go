package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/greybluesea/jwt_mvc_on_aws/database"
	routes "github.com/greybluesea/jwt_mvc_on_aws/routes"
	// "github.com/joho/godotenv"
)

func main() {
	/* err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	*/
	database.ConnectDB()
	engine := html.New("./views", ".html")
	app := fiber.New(
		fiber.Config{
			Views:       engine,
			ViewsLayout: "layout",
		})
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("home", fiber.Map{"Title": "Hello and welcome, this is a JWT-authenticated MVC webapp on AWS👋!"})
		//	return c.SendString("Hello, welcome to the JWT auth GoFiber api 👋!")
	})

	routes.SetAuthRoutes(app)
	routes.SetUserRoutes(app)
	routes.SetSigninRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
