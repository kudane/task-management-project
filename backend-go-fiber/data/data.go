package data

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Priority struct {
	ID   uint   `gorm:"primaryKey;index" json:"id"`
	Name string `json:"name"`
}

type Type struct {
	ID   uint   `gorm:"primaryKey;index" json:"id"`
	Name string `json:"name"`
}

type Task struct {
	ID          uint      `gorm:"primaryKey;index" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"startDate"`
	DueDate     time.Time `json:"dueDate"`
	PriorityID  uint      `json:"priorityId"`
	TypeID      uint      `json:"typeId"`
}

func New() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/sqlite.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Type{}, &Priority{}, &Task{})
	db.Create(&Priority{ID: 1, Name: "Low"})
	db.Create(&Priority{ID: 2, Name: "Medium"})
	db.Create(&Priority{ID: 3, Name: "High"})
	db.Create(&Type{ID: 1, Name: "Personal"})
	db.Create(&Type{ID: 2, Name: "Work"})
	db.Create(&Task{
		ID:          1,
		Name:        "Small",
		Description: "job a",
		StartDate:   time.Now(),
		DueDate:     time.Now(),
		PriorityID:  1,
		TypeID:      1,
	})
	db.Create(&Task{
		ID:          2,
		Name:        "Medium",
		Description: "job b",
		StartDate:   time.Now(),
		DueDate:     time.Now(),
		PriorityID:  2,
		TypeID:      2,
	})
	return db
}
