package model

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	CreateTs  int64          `gorm:"autoCreateTime"`
	UpdateTs  int64          `gorm:"autoCreateTime"`
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
	)
}
