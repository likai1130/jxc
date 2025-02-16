package models

type PurchaseList struct {
	Model
	AmountPayable  float64 `gorm:"column:amount_payable;type:decimal(10,2);not null;comment:'应付金额'" json:"amount_payable"`
	AmountPaid     float64 `gorm:"column:amount_paid;type:decimal(10,2);not null;comment:'已付金额'" json:"amount_paid"`
	PurchaseDate   string  `gorm:"column:purchase_date;type:varchar(100);not null;comment:'采购日期'" json:"purchase_date"`
	PurchaseNumber string  `gorm:"column:purchase_number;type:varchar(100);comment:'采购单号'" json:"purchase_number"`
	Remarks        string  `gorm:"column:remarks;type:varchar(255);comment:'备注'" json:"remarks"`
	State          string  `gorm:"column:state;type:varchar(50);comment:'状态(已付款/未付款)'" json:"state"`
	StorageStatus  string  `gorm:"column:storage_status;type:varchar(45);comment:'入库状态(已入库/未入库)'" json:"storage_tatus"`
	SupplierId     string  `gorm:"column:supplier_id;type:int;not null;comment:'供货商'" json:"supplier_id"`
	ModelTime
	ControlBy
}

func (PurchaseList) TableName() string {
	return "tb_purchase_list"
}
