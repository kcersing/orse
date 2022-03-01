package models

type OrderAmounts struct {
	CommonFields
	TotalAmount float64 `gorm:"column:total_amount" db:"total_amount" json:"total_amount" form:"total_amount"`
	PayAmount float64 `gorm:"column:pay_amount" db:"pay_amount" json:"pay_amount" form:"pay_amount"`
	FreightAmount float64 `gorm:"column:freight_amount" db:"freight_amount" json:"freight_amount" form:"freight_amount"`
	PromotionAmount float64 `gorm:"column:promotion_amount" db:"promotion_amount" json:"promotion_amount" form:"promotion_amount"`
	IntegrationAmount float64 `gorm:"column:integration_amount" db:"integration_amount" json:"integration_amount" form:"integration_amount"`
	CouponId int64 `gorm:"column:coupon_id" db:"coupon_id" json:"coupon_id" form:"coupon_id"`
	CouponAmount float64 `gorm:"column:coupon_amount" db:"coupon_amount" json:"coupon_amount" form:"coupon_amount"`
	DiscountAmount float64 `gorm:"column:discount_amount" db:"discount_amount" json:"discount_amount" form:"discount_amount"`
	OrderId int64 `gorm:"column:order_id" db:"order_id" json:"order_id" form:"order_id"`
}
