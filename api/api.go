package api

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	app.Get("/person/:id", getPerson)
}

// swagger:model Person
type Person struct {
	Name string `json:"name"`
}

// getPerson godoc
// @Summary      Show an account
// @Description  get person by ID
// @Tags         accounts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200 {object} Person
// @Failure      400
// @Router       /person/{id} [get]
func getPerson(c *fiber.Ctx) error {
	return c.JSON(&Person{
		Name: "John",
	})
}
