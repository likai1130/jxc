package models

import (
	"go-admin/common/models"
)

type PurchaseListGoods struct {
	models.Model

	GCode          string  `json:"gCode" gorm:"type:varchar(100);comment:商品编码"`
	GName          string  `json:"gName" gorm:"type:varchar(255);comment:商品名称"`
	Specification  string  `json:"specification" gorm:"type:varchar(255);comment:商品规格"`
	Unit           string  `json:"unit" gorm:"type:varchar(30);comment:单位"`
	Num            int64   `json:"num" gorm:"type:tinyint;comment:数量"`
	Price          float64 `json:"price" gorm:"type:decimal(10,2);comment:单价"`
	TotalPrice     float64 `json:"totalPrice" gorm:"type:decimal(10,2);comment:总价"`
	PurchaseListId int64   `json:"purchaseListId" gorm:"type:bigint;comment:采购单"`
	GoodsId        int64   `json:"goodsId" gorm:"type:bigint;comment:商品"`
	SupplierId     int64   `json:"supplierId" gorm:"type:bigint;comment:供货商"`
	models.ModelTime
	models.ControlBy
}

func (PurchaseListGoods) TableName() string {
	return "tb_purchase_list_goods"
}

func (e *PurchaseListGoods) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PurchaseListGoods) GetId() interface{} {
	return e.Id
}
