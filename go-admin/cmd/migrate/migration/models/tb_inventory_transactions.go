package models

// InventoryTransaction represents the inventory transaction record.
type InventoryTransaction struct {
	Model
	TransactionType string `gorm:"type:varchar(10);comment:操作类型" json:"transaction_type"`
	GCode           string `gorm:"column:g_code;type:varchar(100);comment:'商品编码'" json:"g_code"`
	GName           string `gorm:"column:g_name;type:varchar(255);comment:'商品名称'" json:"g_name"`
	Num             string `gorm:"column:num;type:int;not null;comment:'数量'" json:"num"`
	Specification   string `gorm:"type:varchar(30);comment:商品规格" json:"specification"`
	OrderNumber     uint   `gorm:"column:order_number;type:varchar(100);comment:'订单号'" json:"order_number"`
	GoodsId         int    `gorm:"column:goods_id;type:int;not null;comment:'商品ID'" json:"goods_id"`
	OrderId         int    `gorm:"column:order_id;type:int;not null;comment:'订单ID(销售单/采购单)'" json:"order_id"`
	Remarks         string `gorm:"column:remarks;type:varchar(255);comment:'备注'" json:"remarks"`
	ModelTime
	ControlBy
}

func (InventoryTransaction) TableName() string {
	return "tb_inventory_transactions"
}
