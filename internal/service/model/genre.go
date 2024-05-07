package model

import "time"

type Genre struct {
	ID        int       `gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `gorm:"not null;"`
	Books     []Book    `gorm:"foreignKey:GenreID"`
	CreatedAt time.Time `gorm:"column:created_at;" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"column:updated_at;" json:"updated_at,omitempty"`
}
