package dto

import (
	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GoodsGetPageReq struct {
	dto.Pagination `search:"-"`
	GCode          string `form:"gCode"  search:"type:exact;column:g_code;table:tb_goods" comment:"商品编码"`
	GName          string `form:"gName"  search:"type:contains;column:g_name;table:tb_goods" comment:"商品名称"`
	Producer       string `form:"producer"  search:"type:contains;column:producer;table:tb_goods" comment:"生产厂商"`
	GoodsOrder
}

type GoodsOrder struct {
	Id                  string `form:"idOrder"  search:"type:order;column:id;table:tb_goods"`
	GCode               string `form:"gCodeOrder"  search:"type:order;column:g_code;table:tb_goods"`
	GName               string `form:"gNameOrder"  search:"type:order;column:g_name;table:tb_goods"`
	Specification       string `form:"specificationOrder"  search:"type:order;column:specification;table:tb_goods"`
	Unit                string `form:"unitOrder"  search:"type:order;column:unit;table:tb_goods"`
	PurchasingPrice     string `form:"purchasingPriceOrder"  search:"type:order;column:purchasing_price;table:tb_goods"`
	SellingPrice        string `form:"sellingPriceOrder"  search:"type:order;column:selling_price;table:tb_goods"`
	StockQuantity       string `form:"stockQuantityOrder"  search:"type:order;column:stock_quantity;table:tb_goods"`
	SafetyStock         string `form:"safetyStockOrder"  search:"type:order;column:safety_stock;table:tb_goods"`
	Description         string `form:"descriptionOrder"  search:"type:order;column:description;table:tb_goods"`
	Producer            string `form:"producerOrder"  search:"type:order;column:producer;table:tb_goods"`
	LastPurchasingPrice string `form:"lastPurchasingPriceOrder"  search:"type:order;column:last_purchasing_price;table:tb_goods"`
	CreatedAt           string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_goods"`
	UpdatedAt           string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_goods"`
	DeletedAt           string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_goods"`
	CreateBy            string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_goods"`
	UpdateBy            string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_goods"`
}

func (m *GoodsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GoodsInsertReq struct {
	Id                  int     `json:"-" comment:"主键编码"` // 主键编码
	GCode               string  `json:"gCode" comment:"商品编码"`
	GName               string  `json:"gName" comment:"商品名称"`
	Specification       string  `json:"specification" comment:"商品规格"`
	Unit                string  `json:"unit" comment:"单位"`
	PurchasingPrice     float64 `json:"purchasingPrice" comment:"采购价"`
	SellingPrice        float64 `json:"sellingPrice" comment:"销售指导价"`
	StockQuantity       int64   `json:"stockQuantity" comment:"当前库存量"`
	SafetyStock         int64   `json:"safetyStock" comment:"安全库存量"`
	Description         string  `json:"description" comment:"商品描述"`
	Producer            string  `json:"producer" comment:"生产厂商"`
	LastPurchasingPrice float64 `json:"lastPurchasingPrice" comment:"上次采购价"`
	common.ControlBy
}

func (s *GoodsInsertReq) Generate(model *models.Goods) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GCode = s.GCode
	model.GName = s.GName
	model.Specification = s.Specification
	model.Unit = s.Unit
	model.PurchasingPrice = s.PurchasingPrice
	model.SellingPrice = s.SellingPrice
	model.StockQuantity = s.StockQuantity
	model.SafetyStock = s.SafetyStock
	model.Description = s.Description
	model.Producer = s.Producer
	model.LastPurchasingPrice = s.PurchasingPrice
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *GoodsInsertReq) GetId() interface{} {
	return s.Id
}

type GoodsUpdateReq struct {
	Id                  int     `uri:"id" comment:"主键编码"` // 主键编码
	GCode               string  `json:"gCode" comment:"商品编码"`
	GName               string  `json:"gName" comment:"商品名称"`
	Specification       string  `json:"specification" comment:"商品规格"`
	Unit                string  `json:"unit" comment:"单位"`
	PurchasingPrice     float64 `json:"purchasingPrice" comment:"采购价"`
	SellingPrice        float64 `json:"sellingPrice" comment:"销售指导价"`
	StockQuantity       int64   `json:"stockQuantity" comment:"当前库存量"`
	SafetyStock         int64   `json:"safetyStock" comment:"安全库存量"`
	Description         string  `json:"description" comment:"商品描述"`
	Producer            string  `json:"producer" comment:"生产厂商"`
	LastPurchasingPrice float64 `json:"lastPurchasingPrice" comment:"上次采购价"`
	common.ControlBy
}

func (s *GoodsUpdateReq) Generate(model *models.Goods) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GCode = s.GCode
	model.GName = s.GName
	model.Specification = s.Specification
	model.Unit = s.Unit
	model.PurchasingPrice = s.PurchasingPrice
	model.SellingPrice = s.SellingPrice
	model.StockQuantity = s.StockQuantity
	model.SafetyStock = s.SafetyStock
	model.Description = s.Description
	model.Producer = s.Producer
	model.LastPurchasingPrice = s.PurchasingPrice
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *GoodsUpdateReq) GetId() interface{} {
	return s.Id
}

// GoodsGetReq 功能获取请求参数
type GoodsGetReq struct {
	Id int `uri:"id"`
}

func (s *GoodsGetReq) GetId() interface{} {
	return s.Id
}

// GoodsDeleteReq 功能删除请求参数
type GoodsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GoodsDeleteReq) GetId() interface{} {
	return s.Ids
}
