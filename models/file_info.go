package models

import (
	"time"

	"github.com/google/uuid"
)

type FileInfo struct {
	ID       uuid.UUID `gorm:"default:gen_random_uuid();not null;type:uuid;primary_key;"`
	FileName string    `gorm:"not null"`
	FileType string    `gorm:"not null"`
	FileUrl  string    `gorm:"default:null"`
	FilePath string    `gorm:"default:null"`
	Activate bool      `gorm:"default:true;not null"`
	CreateBy string    `gorm:"not null"`
	UpdateBy string    `gorm:"not null"`
	CreateAt time.Time `gorm:"default:now();not null;"`
	UpdateAt time.Time `gorm:"default:now();not null;"`
}
