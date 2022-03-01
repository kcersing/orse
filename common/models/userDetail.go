package models

type UserDetail struct {
	CommonFields
	UserId int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Age int64 `gorm:"column:age" db:"age" json:"age" form:"age"`
	Rank float64 `gorm:"column:rank" db:"rank" json:"rank" form:"rank"`
	Pic string `gorm:"column:pic" db:"pic" json:"pic" form:"pic"`

}