package model

import "time"

type Author struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Firstname string    `gorm:"not null;"`
	Lastname  string    `gorm:"not null;"`
	Books     []Book    `gorm:"many2many:author_books;"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}
