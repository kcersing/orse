package models

func (Order) TableName() string {
	return "order"
}

// Order 订单信息
type Order struct {
	CommonFields
	UserId      int64  `gorm:"column:user_id;not null;index:order_user_id;comment:会员ID" db:"user_id" json:"user_id" form:"user_id"`
	Sn          string `gorm:"column:sn;not null;collate:utf8mb4_bin;comment:订单编号" db:"sn" json:"sn" form:"sn"`
	Source      string `gorm:"column:source;not null;collate:utf8mb4_bin;comment:订单来源" db:"source" json:"source" form:"source"`
	Status      int8   `gorm:"column:states;not null;default:0;comment:订单状态" db:"states" json:"states" form:"states"`
	Integration int64  `gorm:"column:integration;not null;comment:获得积分" db:"integration" json:"integration" form:"integration"`
	PaymentTime int64  `gorm:"column:payment_time;comment:支付时间" db:"payment_time" json:"payment_time" form:"payment_time"`
	CommentTime int64  `gorm:"column:comment_time;comment:评论时间" db:"comment_time" json:"comment_time" form:"comment_time"`
	Note        string `gorm:"column:note;not null;collate:utf8mb4_bin;comment:备注" db:"note" json:"note" form:"note"`
	Delete      string `gorm:"column:delete;collate:utf8mb4_bin;comment:删除标识" db:"delete" json:"delete" form:"delete"`
}

func (OrderAmounts) TableName() string {
	return "order_amounts"
}

// OrderAmounts 订单金额信息
type OrderAmounts struct {
	CommonFields
	TotalAmount       float64 `gorm:"column:total_amount;not null;default:0;comment:总金额" db:"total_amount" json:"total_amount" form:"total_amount"`
	PayAmount         float64 `gorm:"column:pay_amount;not null;default:0;comment:支付金额" db:"pay_amount" json:"pay_amount" form:"pay_amount"`
	FreightAmount     float64 `gorm:"column:freight_amount;not null;default:0;comment:运费金额" db:"freight_amount" json:"freight_amount" form:"freight_amount"`
	PromotionAmount   float64 `gorm:"column:promotion_amount;not null;default:0;comment:促销金额" db:"promotion_amount" json:"promotion_amount" form:"promotion_amount"`
	IntegrationAmount float64 `gorm:"column:integration_amount;not null;default:0;comment:积分抵扣金额" db:"integration_amount" json:"integration_amount" form:"integration_amount"`
	CouponId          int64   `gorm:"column:coupon_id;index:order_amounts_coupon_id;not null;default:0;comment:优惠卷ID" db:"coupon_id" json:"coupon_id" form:"coupon_id"`
	CouponAmount      float64 `gorm:"column:coupon_amount;not null;default:0;comment:优惠卷优惠金额" db:"coupon_amount" json:"coupon_amount" form:"coupon_amount"`
	DiscountAmount    float64 `gorm:"column:discount_amount;not null;default:0;comment:折扣金额" db:"discount_amount" json:"discount_amount" form:"discount_amount"`
	OrderId           int64   `gorm:"column:order_id;;not null;index:order_amounts_order_id;comment:订单ID" db:"order_id" json:"order_id" form:"order_id"`
}

func (OrderItem) TableName() string {
	return "order_item"
}

// OrderItem 订单商品信息
type OrderItem struct {
	CommonFields
	ProductCateId    int64   `gorm:"column:product_cate_id;not null;index:order_item_product_cate_id;comment:分类ID" db:"product_cate_id" json:"product_cate_id" form:"product_cate_id"`
	ProductId        int64   `gorm:"column:product_id;not null;index:order_item_product_id;comment:商品ID" db:"product_id" json:"product_id" form:"product_id"`
	ProductPic       string  `gorm:"column:product_pic" db:"product_pic" json:"product_pic;comment:商品图片" form:"product_pic"`
	ProductName      string  `gorm:"column:product_name;not null;comment:商品名称" db:"product_name" json:"product_name" form:"product_name"`
	ProductSn        int64   `gorm:"column:product_sn;not null;comment:商品编号" db:"product_sn" json:"product_sn" form:"product_sn"`
	ProductPrice     float64 `gorm:"column:product_price;not null;default:0;comment:商品金额" db:"product_price" json:"product_price" form:"product_price"`
	Quantity         int64   `gorm:"column:quantity;not null;default:1;comment:数量" db:"quantity" json:"quantity" form:"quantity"`
	ProductSpecsId   int64   `gorm:"column:product_specs_id;not null;index:order_item_product_specs_id;comment:商品规格ID" db:"product_specs_id" json:"product_specs_id" form:"product_specs_id"`
	ProductSpecsSn   string  `gorm:"column:product_specs_sn;not null;index:order_item_product_specs_sn;comment:商品规格编号" db:"product_specs_sn" json:"product_specs_sn" form:"product_specs_sn"`
	ProductSpecsAttr string  `gorm:"column:product_specs_attr;not null;" db:"product_specs_attr;comment:商品规格" json:"product_specs_attr" form:"product_specs_attr"`
	OrderId          int64   `gorm:"column:order_id;not null;index:order_item_order_id;comment:订单ID" db:"order_id" json:"order_id" form:"order_id"`
}

func (OrderPay) TableName() string {
	return "order_pay"
}

// OrderPay 订单支付信息
type OrderPay struct {
	CommonFields
	UserId  int64   `gorm:"column:user_id;not null;index:order_pay_user_id;comment:会员ID" db:"user_id" json:"user_id" form:"user_id"`
	Sn      string  `gorm:"column:sn;not null;comment:支付信息编号" db:"sn" json:"sn" form:"sn"`
	Price   float64 `gorm:"column:price;not null;default:0;comment:支付金额" db:"price" json:"price" form:"price"`
	PayMode string  `gorm:"column:pay_mode;not null;comment:支付类型" db:"pay_mode" json:"pay_mode" form:"pay_mode"`
	OrderId int64   `gorm:"column:order_id;not null;index:order_pay_order_id;comment:订单ID" db:"order_id" json:"order_id" form:"order_id"`
}

func (OrderSetting) TableName() string {
	return "order_setting"
}

// OrderSetting 订单流转时间
type OrderSetting struct {
	CommonFields
	OrderOvertime   int64 `gorm:"column:order_overtime;comment:订单关闭时间" db:"order_overtime" json:"order_overtime" form:"order_overtime"`
	ConfirmOvertime int64 `gorm:"column:confirm_overtime;comment:订单确定结束时间" db:"confirm_overtime" json:"confirm_overtime" form:"confirm_overtime"`
	FinishOvertime  int64 `gorm:"column:finish_overtime;comment:订单关闭时间" db:"finish_overtime" json:"finish_overtime" form:"finish_overtime"`
	CommentOvertime int64 `gorm:"column:comment_overtime;comment:评论结束时间" db:"comment_overtime" json:"comment_overtime" form:"comment_overtime"`
	OrderId         int64 `gorm:"column:order_id;not null;index:order_setting_order_id;comment:订单ID" db:"order_id" json:"order_id" form:"order_id"`
}
