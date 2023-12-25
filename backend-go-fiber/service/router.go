package service

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Service struct {
	db *gorm.DB
}

func New(db *gorm.DB) *Service {
	return &Service{db}
}

func (a *Service) Registor(b *fiber.App) {
	list := b.Group("/list")
	{
		list.Get("/task", a.ListTask)
		list.Get("/type", a.ListType)
		list.Get("/priority", a.ListPriorty)
	}

	get := b.Group("/get")
	{
		get.Get("/task/:taskId", a.ByIdTask)
	}

	new := b.Group("/new")
	{
		new.Post("/task", a.NewTask)
	}

	edit := b.Group("/edit")
	{
		edit.Put("/task", a.EditTask)
	}

	delete := b.Group("/delete")
	{
		delete.Delete("/task/:taskId", a.DeleteTask)
	}
}
