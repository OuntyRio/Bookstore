package model

import "time"

type Book struct {
	ID        int    `gorm:"column:id;primaryKey;autoIncrement"`
	Title     string `gorm:"not null;"`
	GenreID   int
	Authors   []Author  `gorm:"many2many:author_books;"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}
