package models

import (
	"time"

	"gorm.io/gorm"
)

type Teacher struct {
	Email     string     `gorm:"primaryKey"`
	Students  []*Student `gorm:"many2many:teacher_student;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
