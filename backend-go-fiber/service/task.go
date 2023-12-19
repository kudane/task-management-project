package service

import (
	"backend/go-fiber/data"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Service) ByIdTask(f *fiber.Ctx) error {
	id := f.Params("taskId")
	var result data.Task
	if err := s.db.First(&result, "ID = ?", id); err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	return f.Status(http.StatusOK).JSON(result)
}

func (s *Service) ListTask(f *fiber.Ctx) error {
	var results []data.Task
	if err := s.db.Find(&results); err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	response := fiber.Map{"items": results}
	return f.Status(http.StatusOK).JSON(response)
}

func (s *Service) NewTask(f *fiber.Ctx) error {
	task := new(data.Task)
	if err := f.BodyParser(task); err != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := s.db.Save(&task); err != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	return f.SendStatus(http.StatusOK)
}

func (s *Service) EditTask(f *fiber.Ctx) error {
	task := new(data.Task)
	if err := f.BodyParser(task); err != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error())
	}
	if err := s.db.Save(&task); err != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	return f.SendStatus(http.StatusOK)
}

func (s *Service) DeleteTask(f *fiber.Ctx) error {
	id := f.Params("taskId")
	if err := s.db.Where("ID = ?", id).Delete(&data.Task{}); err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	return f.SendStatus(http.StatusOK)
}
