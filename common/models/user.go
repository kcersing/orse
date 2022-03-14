package models

func (User) TableName() string {
	return "user"
}

type User struct {
	CommonFields
	Username string `gorm:"column:username;collate:utf8mb4_bin;comment:会员账户" db:"username" json:"username" form:"username"`
	Mobile   string `gorm:"column:mobile;collate:utf8mb4_bin;comment:会员手机号" db:"mobile" json:"mobile" form:"mobile"`
	Pass     string `gorm:"column:pass;collate:utf8mb4_bin;comment:密码" db:"pass" json:"pass" form:"pass"`
	Active   int8   `gorm:"column:active;comment:头像" db:"active" json:"active" form:"active"`
	Status   int8   `gorm:"column:states;comment:状态[0:禁用 1:启用]" db:"states" json:"states" form:"states"`
	Login    int64  `gorm:"column:login;comment:登录时间" db:"login" json:"login" form:"login"`
	Role     int64  `gorm:"column:role;comment:角色ID" db:"role" json:"role" form:"role"`
}

func (UserDetail) TableName() string {
	return "user_detail"
}

type UserDetail struct {
	CommonFields
	UserId int64   `gorm:"column:user_id;not null;index:user_detail_user_id" db:"user_id" json:"user_id" form:"user_id"`
	Name   string  `gorm:"column:name;collate:utf8mb4_bin;comment:姓名" db:"name" json:"name" form:"name"`
	Age    int64   `gorm:"column:age;not null;default:0;comment:年龄" db:"age" json:"age" form:"age"`
	Rank   float64 `gorm:"column:rank;not null;default:0;comment:积分" db:"rank" json:"rank" form:"rank"`
	Pic    string  `gorm:"column:pic;comment:头像" db:"pic" json:"pic" form:"pic"`
}
