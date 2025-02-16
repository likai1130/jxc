package models

type Goods struct {
	Model
	GCode               string  `gorm:"column:g_code;type:varchar(30);comment:商品编码" json:"g_code"`                                  // 商品编码
	GName               string  `gorm:"column:g_name;type:varchar(200);not null;comment:商品名称" json:"g_name"`                        // 商品名称
	Specification       string  `gorm:"column:specification;type:varchar(30);not null;comment:商品规格" json:"specification"`           // 商品规格
	Unit                string  `gorm:"column:unit;type:varchar(30);not null;comment:单位" json:"unit"`                               // 单位
	PurchasingPrice     float64 `gorm:"column:purchasing_price;type:DECIMAL(10,2);not null;comment:采购价" json:"purchasing_price"`    // 采购价
	SellingPrice        float64 `gorm:"column:selling_price;type:DECIMAL(10,2);not null;comment:销售指导价" json:"selling_price"`        // 销售指导价
	StockQuantity       int     `gorm:"column:stock_quantity;type:int;comment:当前库存量" json:"stock_quantity"`                         // 当前库存量
	SafetyStock         int     `gorm:"column:safety_stock;type:int;comment:安全库存量" json:"safety_stock"`                             // 安全库存量
	Description         string  `gorm:"column:description;type:varchar(1000);comment:商品描述" json:"description"`                      // 商品描述
	Producer            string  `gorm:"column:producer;type:varchar(100);comment:生产厂商" json:"producer"`                             // 生产厂商
	LastPurchasingPrice float64 `gorm:"column:last_purchasing_price;type:DECIMAL(10,2);comment:上次采购价" json:"last_purchasing_price"` // 上次采购价
	ModelTime
	ControlBy
}

func (Goods) TableName() string {
	return "tb_goods"
}
