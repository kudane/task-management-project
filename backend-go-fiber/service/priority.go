package service

import (
	"backend/go-fiber/data"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ListPriorty godoc
// @Success 200 {object} []data.Priority
// @Router /list/priority [get]
func (a *Service) ListPriorty(f *fiber.Ctx) error {
	var results []data.Priority
	err := a.db.Find(&results)
	if err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	response := fiber.Map{"items": results}
	return f.Status(http.StatusOK).JSON(response)
}
