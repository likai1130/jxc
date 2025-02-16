package models

import (
	"go-admin/common/models"
)

type SaleList struct {
	models.Model

	AmountPayable  float64 `json:"amountPayable" gorm:"type:decimal(10,2);comment:应付金额"`
	AmountPaid     float64 `json:"amountPaid" gorm:"type:decimal(10,2);comment:已付金额"`
	SaleDate       string  `json:"saleDate" gorm:"type:varchar(100);comment:销售日期"`
	SaleNumber     string  `json:"saleNumber" gorm:"type:varchar(100);comment:销售单号"`
	Remarks        string  `json:"remarks" gorm:"type:varchar(255);comment:备注"`
	State          string  `json:"state" gorm:"type:varchar(50);comment:状态"`
	ShipmentStatus string  `gorm:"type:varchar(45);comment:'出库状态(已出库/未出库)'" json:"shipmentStatus"`
	CustomerId     int64   `json:"customerId" gorm:"type:tinyint;comment:消费者ID"`
	models.ModelTime
	models.ControlBy
}

func (SaleList) TableName() string {
	return "tb_sale_list"
}

func (e *SaleList) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *SaleList) GetId() interface{} {
	return e.Id
}
