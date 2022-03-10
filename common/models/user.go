package models
func (User) TableName() string {
	return "user"
}
type User struct {
	CommonFields
	Username string `gorm:"column:username" db:"username" json:"username" form:"username"`
	Mobile   string `gorm:"column:mobile" db:"mobile" json:"mobile" form:"mobile"`
	Pass     string `gorm:"column:pass" db:"pass" json:"pass" form:"pass"`
	Active   int8   `gorm:"column:active" db:"active" json:"active" form:"active"`
	State    string `gorm:"column:state" db:"state" json:"state" form:"state"`
	Login    int64  `gorm:"column:login" db:"login" json:"login" form:"login"`
	Role     int64  `gorm:"column:role" db:"role" json:"role" form:"role"`
}
func (UserDetail) TableName() string {
	return "user_detail"
}
type UserDetail struct {
	CommonFields
	UserId int64   `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Name   string  `gorm:"column:name" db:"name" json:"name" form:"name"`
	Age    int64   `gorm:"column:age" db:"age" json:"age" form:"age"`
	Rank   float64 `gorm:"column:rank" db:"rank" json:"rank" form:"rank"`
	Pic    string  `gorm:"column:pic" db:"pic" json:"pic" form:"pic"`
}
