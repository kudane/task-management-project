package service

import (
	"backend/go-fiber/data"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// ByIdTask godoc
// @Success 200 {object} data.Task
// @Router /list/task/:taskId [get]
func (s *Service) ByIdTask(f *fiber.Ctx) error {
	id := f.Params("taskId")
	var result data.Task
	if err := s.db.First(&result, "ID = ?", id); err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	return f.Status(http.StatusOK).JSON(result)
}

// ListTask godoc
// @Success 200 {object} []data.Task
// @Router /list/task [get]
func (s *Service) ListTask(f *fiber.Ctx) error {
	var results []data.Task
	query := s.db.Model(&results)
	if q := f.QueryInt("type"); q != 0 {
		query = query.Where(&data.Task{TypeID: uint(q)})
	}
	if q := f.QueryInt("priority"); q != 0 {
		query = query.Where(&data.Task{PriorityID: uint(q)})
	}
	if err := query.Find(&results); err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	response := fiber.Map{"items": results}
	return f.Status(http.StatusOK).JSON(response)
}

// NewTask godoc
// @Params data.Task
// @Success 200 {object} []data.Task
// @Router /new/task [post]
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

// EditTask godoc
// @Params data.Task
// @Success 200 {object} []data.Task
// @Router /edit/task [put]
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

// NewTask godoc
// @Success 200 {object} []data.Task
// @Router /delete/task/:taskId [put]
func (s *Service) DeleteTask(f *fiber.Ctx) error {
	id := f.Params("taskId")
	if err := s.db.Where("ID = ?", id).Delete(&data.Task{}); err.Error != nil {
		return f.Status(http.StatusBadRequest).SendString(err.Error.Error())
	}
	return f.SendStatus(http.StatusOK)
}
