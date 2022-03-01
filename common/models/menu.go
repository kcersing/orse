package models

type Menu struct {
	CommonFields
	Title string `gorm:"column:title" db:"title" json:"title" form:"title"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Url string `gorm:"column:url" db:"url" json:"url" form:"url"`
	ParentId int64 `gorm:"column:parent_id" db:"parent_id" json:"parent_id" form:"parent_id"`
	Component string `gorm:"column:component" db:"component" json:"component" form:"component"`
	Level int64 `gorm:"column:level" db:"level" json:"level" form:"level"`
	Sort int64 `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`
	Status string `gorm:"column:status" db:"status" json:"status" form:"status"`
	Icon string `gorm:"column:icon" db:"icon" json:"icon" form:"icon"`
	Hidden string `gorm:"column:hidden" db:"hidden" json:"hidden" form:"hidden"`
	Desc string `gorm:"column:desc" db:"desc" json:"desc" form:"desc"`

}