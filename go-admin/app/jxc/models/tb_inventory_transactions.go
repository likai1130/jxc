package models

import (
	"go-admin/common/models"
)

type InventoryTransactions struct {
	models.Model

	TransactionType string `json:"transactionType" gorm:"type:varchar(10);comment:操作类型"`
	GCode           string `json:"gCode" gorm:"type:varchar(100);comment:商品编码"`
	GName           string `json:"gName" gorm:"type:varchar(255);comment:商品名称"`
	Specification   string ` json:"specification" gorm:"type:varchar(30);comment:商品规格"`
	Num             int64  `json:"num" gorm:"type:tinyint;comment:数量"`
	OrderNumber     string `json:"orderNumber" gorm:"type:varchar(100);comment:订单号"`
	GoodsId         int64  `json:"goodsId" gorm:"type:bigint;comment:商品ID"`
	OrderId         int64  `json:"orderId" gorm:"type:bigint;comment:订单ID"`
	Remarks         string `json:"remarks" gorm:"type:varchar(255);comment:备注"`
	models.ModelTime
	models.ControlBy
}

func (InventoryTransactions) TableName() string {
	return "tb_inventory_transactions"
}

func (e *InventoryTransactions) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *InventoryTransactions) GetId() interface{} {
	return e.Id
}
