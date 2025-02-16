package models

import (
	"go-admin/common/models"
)

type SaleListGoods struct {
	models.Model

	GCode         string  `json:"gCode" gorm:"type:varchar(100);comment:商品编码"`
	GName         string  `json:"gName" gorm:"type:varchar(255);comment:商品名称"`
	Specification string  `json:"specification" gorm:"type:varchar(255);comment:商品规格"`
	Unit          string  `json:"unit" gorm:"type:varchar(30);comment:单位"`
	Num           int64   `json:"num" gorm:"type:bigint;comment:数量"`
	Price         float64 `json:"price" gorm:"type:decimal(10,2);comment:采购价"`
	TotalPrice    float64 `json:"totalPrice" gorm:"type:decimal(10,2);comment:总价"`
	SaleListId    int64   `json:"saleListId" gorm:"type:bigint;comment:销售单"`
	GoodsId       int64   `json:"goodsId" gorm:"type:bigint;comment:商品信息"`
	models.ModelTime
	models.ControlBy
}

func (SaleListGoods) TableName() string {
	return "tb_sale_list_goods"
}

func (e *SaleListGoods) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SaleListGoods) GetId() interface{} {
	return e.Id
}
