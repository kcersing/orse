package models

// Order 订单信息
type Order struct {
	CommonFields
	UserId      int64  `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Sn          string `gorm:"column:sn" db:"sn" json:"sn" form:"sn"`
	Source      string `gorm:"column:source" db:"source" json:"source" form:"source"`
	Status      string `gorm:"column:status" db:"status" json:"status" form:"status"`
	Integration int64  `gorm:"column:integration" db:"integration" json:"integration" form:"integration"`
	PaymentTime int64  `gorm:"column:payment_time" db:"payment_time" json:"payment_time" form:"payment_time"`
	Note        string `gorm:"column:note" db:"note" json:"note" form:"note"`
	CommentTime int64  `gorm:"column:comment_time" db:"comment_time" json:"comment_time" form:"comment_time"`
	Delete      string `gorm:"column:delete" db:"delete" json:"delete" form:"delete"`
}

// OrderAmounts 订单金额信息
type OrderAmounts struct {
	CommonFields
	TotalAmount       float64 `gorm:"column:total_amount" db:"total_amount" json:"total_amount" form:"total_amount"`
	PayAmount         float64 `gorm:"column:pay_amount" db:"pay_amount" json:"pay_amount" form:"pay_amount"`
	FreightAmount     float64 `gorm:"column:freight_amount" db:"freight_amount" json:"freight_amount" form:"freight_amount"`
	PromotionAmount   float64 `gorm:"column:promotion_amount" db:"promotion_amount" json:"promotion_amount" form:"promotion_amount"`
	IntegrationAmount float64 `gorm:"column:integration_amount" db:"integration_amount" json:"integration_amount" form:"integration_amount"`
	CouponId          int64   `gorm:"column:coupon_id" db:"coupon_id" json:"coupon_id" form:"coupon_id"`
	CouponAmount      float64 `gorm:"column:coupon_amount" db:"coupon_amount" json:"coupon_amount" form:"coupon_amount"`
	DiscountAmount    float64 `gorm:"column:discount_amount" db:"discount_amount" json:"discount_amount" form:"discount_amount"`
	OrderId           int64   `gorm:"column:order_id" db:"order_id" json:"order_id" form:"order_id"`
}
// OrderItem 订单商品信息
type OrderItem struct {
	CommonFields
	Sn               string  `gorm:"column:sn" db:"sn" json:"sn" form:"sn"`
	ProductCateId    int64   `gorm:"column:product_cate_id" db:"product_cate_id" json:"product_cate_id" form:"product_cate_id"`
	ProductId        int64   `gorm:"column:product_id" db:"product_id" json:"product_id" form:"product_id"`
	ProductPic       string  `gorm:"column:product_pic" db:"product_pic" json:"product_pic" form:"product_pic"`
	ProductName      string  `gorm:"column:product_name" db:"product_name" json:"product_name" form:"product_name"`
	ProductSn        int64   `gorm:"column:product_sn" db:"product_sn" json:"product_sn" form:"product_sn"`
	ProductPrice     float64 `gorm:"column:product_price" db:"product_price" json:"product_price" form:"product_price"`
	Quantity         int64   `gorm:"column:quantity" db:"quantity" json:"quantity" form:"quantity"`
	ProductSpecsId   int64   `gorm:"column:product_specs_id" db:"product_specs_id" json:"product_specs_id" form:"product_specs_id"`
	ProductSpecsSn   string  `gorm:"column:product_specs_sn" db:"product_specs_sn" json:"product_specs_sn" form:"product_specs_sn"`
	ProductSpecsAttr string  `gorm:"column:product_specs_attr" db:"product_specs_attr" json:"product_specs_attr" form:"product_specs_attr"`
	OrderId          int64   `gorm:"column:order_id" db:"order_id" json:"order_id" form:"order_id"`
}

// OrderPay 订单支付信息
type OrderPay struct {
	CommonFields
	UserId int64 `gorm:"column:user_id" db:"user_id" json:"user_id" form:"user_id"`
	Sn string `gorm:"column:sn" db:"sn" json:"sn" form:"sn"`
	Price float64 `gorm:"column:price" db:"price" json:"price" form:"price"`
	PayMode string `gorm:"column:pay_mode" db:"pay_mode" json:"pay_mode" form:"pay_mode"`
	OrderId int64 `gorm:"column:order_id" db:"order_id" json:"order_id" form:"order_id"`
}

// OrderSetting 订单流转时间
type OrderSetting struct {
	CommonFields
	OrderOvertime int64 `gorm:"column:order_overtime" db:"order_overtime" json:"order_overtime" form:"order_overtime"`
	ConfirmOvertime int64 `gorm:"column:confirm_overtime" db:"confirm_overtime" json:"confirm_overtime" form:"confirm_overtime"`
	FinishOvertime int64 `gorm:"column:finish_overtime" db:"finish_overtime" json:"finish_overtime" form:"finish_overtime"`
	CommentOvertime int64 `gorm:"column:comment_overtime" db:"comment_overtime" json:"comment_overtime" form:"comment_overtime"`
}