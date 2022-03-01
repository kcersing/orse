package models

type User struct {
	CommonFields
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Mobile string `gorm:"column:mobile" db:"mobile" json:"mobile" form:"mobile"`
	Pass string `gorm:"column:pass" db:"pass" json:"pass" form:"pass"`
	Active int8 `gorm:"column:active" db:"active" json:"active" form:"active"`
	State string `gorm:"column:state" db:"state" json:"state" form:"state"`
	Login int64 `gorm:"column:login" db:"login" json:"login" form:"login"`
	Role int64 `gorm:"column:role" db:"role" json:"role" form:"role"`
}

