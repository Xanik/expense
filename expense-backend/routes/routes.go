package routes

import (
	"expense-backend/store"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/save", Save)
	app.Get("/fetch", Fetch)
}

func Save(c *fiber.Ctx) error {
	var expense store.Payload

	err := c.BodyParser(&expense)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	// Store expense
	store.SaveExpense(expense)
	return c.Status(200).JSON(map[string]bool{
		"success": true,
	})
}

func Fetch(c *fiber.Ctx) error {
	var expenses = store.FetchExpenses()
	return c.Status(200).JSON(map[string]interface{}{
		"success": true,
		"data": map[string]interface{}{
			"expenses": expenses,
		},
	})
}
