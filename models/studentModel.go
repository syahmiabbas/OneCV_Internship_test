package models

import (
	"time"

	"gorm.io/gorm"
)

type Student struct {
	Email     string `gorm:"primaryKey"`
	Suspended bool
	Teachers  []*Teacher `gorm:"many2many:teacher_student;"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
