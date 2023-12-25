package service

import (
	"backend/go-fiber/data"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ListType godoc
// @Success 200 {object} []data.Type
// @Router /list/type [get]
func (a *Service) ListType(f *fiber.Ctx) error {
	var results []data.Type
	err := a.db.Find(&results)
	if err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	response := fiber.Map{"items": results}
	return f.Status(http.StatusOK).JSON(response)
}
