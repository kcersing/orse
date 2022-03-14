package models

func (ProductCate) TableName() string {
	return "product_cate"
}

// ProductCate 商品类别信息
type ProductCate struct {
	CommonFields
	ParentId  int64  `gorm:"column:parent_id;index:product_cate_parent_id;not null;default:0;comment:上级ID" db:"parent_id" json:"parent_id" form:"parent_id"`
	Status    int8   `gorm:"column:states;not null;default:0;comment:状态[0:禁用 1:启用]" db:"states" json:"states" form:"states"`
	Name      string `gorm:"column:name;collate:utf8mb4_bin;comment:名称" db:"name" json:"name" form:"name"`
	Sort      int64  `gorm:"column:sort;not null;default:0;comment:排序" db:"sort" json:"sort" form:"sort"`
}

func (Product) TableName() string {
	return "product"
}

// Product 商品信息
type Product struct {
	CommonFields
	Name      string `gorm:"column:name;collate:utf8mb4_bin;comment:名称" db:"name" json:"name" form:"name"`
	Status    int8   `gorm:"column:states;not null;default:0;comment:状态[0:禁用 1:启用]" db:"states" json:"states" form:"states"`
	CreateId  int64  `gorm:"column:create_id;not null;default:0;" db:"create_id" json:"create_id" form:"create_id"`
	CateId    int64  `gorm:"column:cate_id;not null;default:0;index:product_cate_id;comment:分类ID" db:"cate_id" json:"cate_id" form:"cate_id"`
}

func (ProductAttributeKey) TableName() string {
	return "product_attribute_key"
}

// ProductAttributeKey 商品信息
type ProductAttributeKey struct {
	CommonFields
	ProductId int64  `gorm:"column:product_id;not null;default:0;index:product_attribute_key_product_id;comment:商品ID" db:"product_id" json:"product_id" form:"product_id"`
	Name      string `gorm:"column:name;collate:utf8mb4_bin;comment:商品属性名称" db:"name" json:"name" form:"name"`
	Sort      int64  `gorm:"column:sort;not null;default:0;comment:排序" db:"sort" json:"sort" form:"sort"`
}

func (ProductAttributeValue) TableName() string {
	return "product_attribute_value"
}

// ProductAttributeValue 商品信息
type ProductAttributeValue struct {
	CommonFields
	Value     string `gorm:"column:value,collate:utf8mb4_bin;comment:商品属性值" db:"value" json:"value" form:"value"`
	Sort      int64  `gorm:"column:sort;not null;default:0;comment:排序" db:"sort" json:"sort" form:"sort"`
	KeyId     int64  `gorm:"column:key_id;not null;default:0;index:product_attribute_value_key_id;comment:商品属性ID" db:"key_id" json:"key_id" form:"key_id"`
}

func (ProductSpecs) TableName() string {
	return "product_specs"
}

// ProductSpecs 商品信息
type ProductSpecs struct {
	CommonFields
	Name      string  `gorm:"column:name;comment:商品规格名称" db:"name" json:"name" form:"name"`
	Sn        string  `gorm:"column:sn;comment:商品规格编号" db:"sn" json:"sn" form:"sn"`
	Stock     int64   `gorm:"column:stock;not null;default:0;comment:库存" db:"stock" json:"stock" form:"stock"`
	Sales     int64   `gorm:"column:sales;not null;default:0;comment:已销售数量" db:"sales" json:"sales" form:"sales"`
	Price     float64 `gorm:"column:price;not null;default:0;comment:价格" db:"price" json:"price" form:"price"`
	SalePrice float64 `gorm:"column:sale_price;not null;default:0;comment:销售价格" db:"sale_price" json:"sale_price" form:"sale_price"`
	CreateId  int64   `gorm:"column:create_id;not null;default:0;comment:商品规格" db:"create_id" json:"create_id" form:"create_id"`
	ProductId int64   `gorm:"column:product_id;not null;default:0;index:product_specs_product_id;comment:商品ID" db:"product_id" json:"product_id" form:"product_id"`
}

func (ProductSpecsItem) TableName() string {
	return "product_specs_item"
}

// ProductSpecsItem 商品信息
type ProductSpecsItem struct {
	CommonFields
	ValueId        int64 `gorm:"column:value_id;not null;default:0;index:product_specs_item_value_id;comment:商品规格项目值" db:"value_id" json:"value_id" form:"value_id"`
	Sort           int64 `gorm:"column:sort;not null;default:0;comment:排序" db:"sort" json:"sort" form:"sort"`
	ProductSpecsId int64 `gorm:"column:product_specs_id;not null;default:0;index:product_specs_item_product_specs_id;comment:商品规格ID" db:"product_specs_id" json:"product_specs_id" form:"product_specs_id"`
}

func (ProductSpecsItemValues) TableName() string {
	return "product_specs_item_values"
}

// ProductSpecsItemValues 商品信息
type ProductSpecsItemValues struct {
	ProductSpecsItemId      int64 `gorm:"column:product_specs_item_id;not null;default:0;index:product_specs_item_values_product_specs_item_id;comment:商品规格项目ID" db:"product_specs_item_id" json:"product_specs_item_id" form:"product_specs_item_id"`
	ProductAttributeValueId int64 `gorm:"column:product_attribute_value_id;not null;default:0;index:product_specs_item_values_product_attribute_value_id;comment:商品属性值ID" db:"product_attribute_value_id" json:"product_attribute_value_id" form:"product_attribute_value_id"`
}
