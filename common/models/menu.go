package models

func (Menu) TableName() string {
	return "menu"
}

type Menu struct {
	CommonFields
	ParentId  int64  `gorm:"column:parent_id;index:menu_parent_id;comment:上级ID" db:"parent_id" json:"parent_id" form:"parent_id"`
	Title     string `gorm:"column:title;collate:utf8mb4_bin;comment:标题" db:"title" json:"title" form:"title"`
	Name      string `gorm:"column:name;collate:utf8mb4_bin;comment:名称" db:"name" json:"name" form:"name"`
	Url       string `gorm:"column:url;collate:utf8mb4_bin;comment:路由" db:"url" json:"url" form:"url"`
	Component string `gorm:"column:component;collate:utf8mb4_bin;comment:前端组件" db:"component" json:"component" form:"component"`
	Level     int64  `gorm:"column:level;not null;default:0;comment:深度等级" db:"level" json:"level" form:"level"`
	Sort      int64  `gorm:"column:sort;not null;default:0;comment:排序权重" db:"sort" json:"sort" form:"sort"`
	Status    int8   `gorm:"column:states;not null;default:0;comment:状态[0:禁用 1:启用]" db:"states" json:"states" form:"states"`
	Icon      string `gorm:"column:icon;collate:utf8mb4_bin;comment:图标" db:"icon" json:"icon" form:"icon"`
	Hidden    string `gorm:"column:hidden;collate:utf8mb4_bin" db:"hidden" json:"hidden" form:"hidden"`
	Desc      string `gorm:"column:desc;collate:utf8mb4_bin" db:"desc" json:"desc" form:"desc"`
}
