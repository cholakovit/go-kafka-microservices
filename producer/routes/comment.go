package routes

import (
	"encoding/json"
	"producer/kafka"
	"producer/models"

	"github.com/gofiber/fiber/v2"
)

func RegisterCommentRoutes(app fiber.Router, kafkaProducer *kafka.Producer) {
	app.Post("/comments", func (c *fiber.Ctx) error  {
		return createComment(c, kafkaProducer)
	})
}

func createComment(c *fiber.Ctx, kafkaProducer *kafka.Producer) error {
	comment := new(models.Comment)
	if err := c.BodyParser(comment); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": err,
		})
	}

	message, err := json.Marshal(comment)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"message": "Failed to process comment",
		})
	}

	kafkaProducer.PushComment("comments", message)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Comment pushed successfully",
		"comment": comment,
	})
}