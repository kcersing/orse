package models

import (
	"time"
)

type CommonFields struct {
	Id        int64     `gorm:"primaryKey,column:id" db:"id" json:"id" form:"id"`
	CreatedAt time.Time `gorm:"column:created_at;priority:100" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;priority:100" db:"updated_at" json:"updated_at" form:"updated_at"`
}
