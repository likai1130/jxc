package models

type SaleList struct {
	Model
	AmountPayable  float64 `gorm:"column:amount_payable;type:decimal(10,2);not null;comment:'应付金额'" json:"amount_payable"`
	AmountPaid     float64 `gorm:"column:amount_paid;type:decimal(10,2);not null;comment:'已付金额'" json:"amount_paid"`
	SaleDate       string  `gorm:"column:sale_date;type:varchar(100);not null;comment:'销售日期'" json:"sale_date"`
	SaleNumber     string  `gorm:"column:sale_number;type:varchar(100);comment:'销售单号'" json:"sale_number"`
	Remarks        string  `gorm:"column:remarks;type:varchar(255);comment:'备注'" json:"remarks"`
	State          string  `gorm:"column:state;type:varchar(50);comment:'状态(已付款/未付款)'" json:"state"`
	ShipmentStatus string  `gorm:"column:shipment_status;type:varchar(45);comment:'出库状态(已出库/未出库)'" json:"shipment_status"`
	CustomerId     string  `gorm:"column:customer_id;type:int;not null;comment:'消费者ID'" json:"consumer_id"`
	ModelTime
	ControlBy
}

func (SaleList) TableName() string {
	return "tb_sale_list"
}
