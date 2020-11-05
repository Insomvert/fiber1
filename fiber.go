package main
import "fmt"
import "github.com/gofiber/fiber/v2"
func main() {
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
	fmt.Println(`listen at :3000`)
	app.Listen(":3000")
}
