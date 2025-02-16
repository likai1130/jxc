package models

import (
	"go-admin/common/models"
)

type Goods struct {
	models.Model
	GCode               string  `json:"gCode" gorm:"type:varchar(30);comment:商品编码"`
	GName               string  `json:"gName" gorm:"type:varchar(200);comment:商品名称"`
	Specification       string  `json:"specification" gorm:"type:varchar(30);comment:商品规格"`
	Unit                string  `json:"unit" gorm:"type:varchar(30);comment:单位"`
	PurchasingPrice     float64 `json:"purchasingPrice" gorm:"type:decimal(10,2);comment:采购价"`
	SellingPrice        float64 `json:"sellingPrice" gorm:"type:decimal(10,2);comment:销售指导价"`
	StockQuantity       int64   `json:"stockQuantity" gorm:"type:bigint;comment:当前库存量"`
	SafetyStock         int64   `json:"safetyStock" gorm:"type:bigint;comment:安全库存量"`
	Description         string  `json:"description" gorm:"type:varchar(1000);comment:商品描述"`
	Producer            string  `json:"producer" gorm:"type:varchar(100);comment:生产厂商"`
	LastPurchasingPrice float64 `json:"lastPurchasingPrice" gorm:"type:decimal(10,2);comment:上次采购价"`
	SaleNum             int64   `json:"saleNum" gorm:"-"`
	models.ModelTime
	models.ControlBy
}

func (Goods) TableName() string {
	return "tb_goods"
}

func (e *Goods) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Goods) GetId() interface{} {
	return e.Id
}
