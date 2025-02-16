package models

type GoodsUnit struct {
	Model
	Name string `gorm:"type:varchar(10);comment:商品单位" json:"Name"` // 单位名称
	ModelTime
	ControlBy
}

func (GoodsUnit) TableName() string {
	return "tb_goods_unit"
}
