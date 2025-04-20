package models

import (
	"go-admin/common/models"
)

type PurchaseList struct {
	models.Model

	AmountPayable  float64 `json:"amountPayable" gorm:"type:decimal(10,2);comment:应付金额"`
	AmountPaid     float64 `json:"amountPaid" gorm:"type:decimal(10,2);comment:已付金额"`
	PurchaseDate   string  `json:"purchaseDate" gorm:"type:varchar(100);comment:采购日期"`
	PurchaseNumber string  `json:"purchaseNumber" gorm:"type:varchar(100);comment:采购单号"`
	Remarks        string  `json:"remarks" gorm:"type:varchar(255);comment:备注"`
	State          string  `json:"state" gorm:"type:varchar(50);comment:是否付款"`
	StorageStatus  string  `json:"storageStatus" gorm:"type:varchar(45);comment:是否入库"`
	SupplierId     int64   `json:"supplierId" gorm:"type:tinyint;comment:供货商"`
	SaleNumber     string  `json:"saleNumber" gorm:"type:varchar(100);comment:销售单号"`
	models.ModelTime
	models.ControlBy
}

func (PurchaseList) TableName() string {
	return "tb_purchase_list"
}

func (e *PurchaseList) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *PurchaseList) GetId() interface{} {
	return e.Id
}
