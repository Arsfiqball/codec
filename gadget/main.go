package gadget

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
}

type Gadget struct {
}

func (g *Gadget) FiberApp() *fiber.App {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	return app
}

func (g *Gadget) WatermillSubscriber(pub message.Publisher, sub message.Subscriber, router *message.Router) {
	router.AddNoPublisherHandler("feature", "feature", sub, func(msg *message.Message) error {
		return nil
	})
}
