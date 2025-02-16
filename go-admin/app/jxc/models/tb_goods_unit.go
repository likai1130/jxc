package models

import (

	"go-admin/common/models"

)

type GoodsUnit struct {
    models.Model
    
    Name string `json:"name" gorm:"type:varchar(10);comment:商品单位"` 
    models.ModelTime
    models.ControlBy
}

func (GoodsUnit) TableName() string {
    return "tb_goods_unit"
}

func (e *GoodsUnit) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *GoodsUnit) GetId() interface{} {
	return e.Id
}