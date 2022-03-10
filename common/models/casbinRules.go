package models

type CasbinRules struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	Ptype string `gorm:"column:ptype" db:"ptype" json:"ptype" form:"ptype"`
	Role string `gorm:"column:v0" db:"v0" json:"v0" form:"v0"`
	Path string `gorm:"column:v1" db:"v1" json:"v1" form:"v1"`
	Method string `gorm:"column:v2" db:"v2" json:"v2" form:"v2"`
}