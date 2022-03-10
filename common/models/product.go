package models

// ProductCate 商品类别信息
type ProductCate struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	ParentId int64 `gorm:"column:parent_id" db:"parent_id" json:"parent_id" form:"parent_id"`
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Sort int64 `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`
}
// Product 商品信息
type Product struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Status string `gorm:"column:status" db:"status" json:"status" form:"status"`
	CreateId int64 `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`
	CateId int64 `gorm:"column:cate_id" db:"cate_id" json:"cate_id" form:"cate_id"`
}
// ProductAttributeKey 商品信息
type ProductAttributeKey struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	ProductId int64 `gorm:"column:product_id" db:"product_id" json:"product_id" form:"product_id"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Sort int64 `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`
}
// ProductAttributeValue 商品信息
type ProductAttributeValue struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	Value string `gorm:"column:value" db:"value" json:"value" form:"value"`
	Sort int64 `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`
	KeyId int64 `gorm:"column:key_id" db:"key_id" json:"key_id" form:"key_id"`
}
// ProductSpecs 商品信息
type ProductSpecs struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	Name string `gorm:"column:name" db:"name" json:"name" form:"name"`
	Sn string `gorm:"column:sn" db:"sn" json:"sn" form:"sn"`
	Stock int64 `gorm:"column:stock" db:"stock" json:"stock" form:"stock"`
	Sales int64 `gorm:"column:sales" db:"sales" json:"sales" form:"sales"`
	Price float64 `gorm:"column:price" db:"price" json:"price" form:"price"`
	SalePrice float64 `gorm:"column:sale_price" db:"sale_price" json:"sale_price" form:"sale_price"`
	CreateId int64 `gorm:"column:create_id" db:"create_id" json:"create_id" form:"create_id"`
	ProductId int64 `gorm:"column:product_id" db:"product_id" json:"product_id" form:"product_id"`
}
// ProductSpecsItem 商品信息
type ProductSpecsItem struct {
	Id int64 `gorm:"column:id" db:"id" json:"id" form:"id"`
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	ValueId int64 `gorm:"column:value_id" db:"value_id" json:"value_id" form:"value_id"`
	Sort int64 `gorm:"column:sort" db:"sort" json:"sort" form:"sort"`
	ProductSpecsId int64 `gorm:"column:product_specs_id" db:"product_specs_id" json:"product_specs_id" form:"product_specs_id"`
}

// ProductSpecsItemValues 商品信息
type ProductSpecsItemValues struct {
	ProductSpecsItemId int64 `gorm:"column:product_specs_item_id" db:"product_specs_item_id" json:"product_specs_item_id" form:"product_specs_item_id"`
	ProductAttributeValueId int64 `gorm:"column:product_attribute_value_id" db:"product_attribute_value_id" json:"product_attribute_value_id" form:"product_attribute_value_id"`
}
