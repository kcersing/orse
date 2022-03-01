package models
type Order struct {
	CommonFields
	UserId int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Sn string `gorm:"column:sn" db:"sn" json:"sn" form:"sn"`
	Source string `gorm:"column:source" db:"source" json:"source" form:"source"`
	Status string `gorm:"column:status" db:"status" json:"status" form:"status"`
	Integration int64 `gorm:"column:integration" db:"integration" json:"integration" form:"integration"`
	PaymentTime int64 `gorm:"column:payment_time" db:"payment_time" json:"payment_time" form:"payment_time"`
	Note string `gorm:"column:note" db:"note" json:"note" form:"note"`
	CommentTime int64 `gorm:"column:comment_time" db:"comment_time" json:"comment_time" form:"comment_time"`
	Delete string `gorm:"column:delete" db:"delete" json:"delete" form:"delete"`
}