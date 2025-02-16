package models

type PurchaseSaleListGoods struct {
	Model
	GCode          string  `gorm:"column:g_code;type:varchar(100);comment:'商品编码'" json:"g_code"`
	GName          string  `gorm:"column:g_name;type:varchar(255);comment:'商品名称'" json:"g_name"`
	Specification  string  `gorm:"column:specification;type:varchar(255);comment:'商品规格'" json:"specification"`
	Unit           string  `gorm:"column:unit;type:varchar(30);not null;comment:单位" json:"unit"` // 单位
	Num            string  `gorm:"column:num;type:int;not null;comment:'数量'" json:"num"`
	Price          float64 `gorm:"column:price;type:decimal(10,2);not null;comment:'单价'" json:"price"`
	TotalPrice     float64 `gorm:"column:total_price;type:decimal(10,2);not null;comment:'总价'" json:"total_price"`
	PurchaseListId int     `gorm:"column:purchase_list_id;type:int;not null;comment:'采购单'" json:"purchase_list_id"`
	GoodsId        int     `gorm:"column:goods_id;type:int;not null;comment:'商品'" json:"goods_id"`
	SupplierId     int     `gorm:"column:supplier_id;type:int;not null;comment:'供货商'" json:"supplier_id"`
	ModelTime
	ControlBy
}

func (PurchaseSaleListGoods) TableName() string {
	return "tb_purchase_list_goods"
}
