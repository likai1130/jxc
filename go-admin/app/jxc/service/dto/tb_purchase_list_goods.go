package dto

import (
	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PurchaseListGoodsGetPageReq struct {
	dto.Pagination `search:"-"`
	GName          string `form:"gName"  search:"type:contains;column:g_name;table:tb_purchase_list_goods" comment:"商品名称"`
	PurchaseListId string `form:"purchaseListId"  search:"type:exact;column:purchase_list_id;table:tb_purchase_list_goods" comment:"采购单"`
	PurchaseListGoodsOrder
}

type PurchaseListGoodsOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:tb_purchase_list_goods"`
	GCode          string `form:"gCodeOrder"  search:"type:order;column:g_code;table:tb_purchase_list_goods"`
	GName          string `form:"gNameOrder"  search:"type:order;column:g_name;table:tb_purchase_list_goods"`
	Specification  string `form:"specificationOrder"  search:"type:order;column:specification;table:tb_purchase_list_goods"`
	Unit           string `form:"unitOrder"  search:"type:order;column:unit;table:tb_purchase_list_goods"`
	Num            string `form:"numOrder"  search:"type:order;column:num;table:tb_purchase_list_goods"`
	Price          string `form:"priceOrder"  search:"type:order;column:price;table:tb_purchase_list_goods"`
	TotalPrice     string `form:"totalPriceOrder"  search:"type:order;column:total_price;table:tb_purchase_list_goods"`
	PurchaseListId string `form:"purchaseListIdOrder"  search:"type:order;column:purchase_list_id;table:tb_purchase_list_goods"`
	GoodsId        string `form:"goodsIdOrder"  search:"type:order;column:goods_id;table:tb_purchase_list_goods"`
	SupplierId     string `form:"supplierIdOrder"  search:"type:order;column:supplier_id;table:tb_purchase_list_goods"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_purchase_list_goods"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_purchase_list_goods"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_purchase_list_goods"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_purchase_list_goods"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_purchase_list_goods"`
}

func (m *PurchaseListGoodsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PurchaseListGoodsInsertReq struct {
	Id             int     `json:"-" comment:"主键编码"` // 主键编码
	GCode          string  `json:"gCode" comment:"商品编码"`
	GName          string  `json:"gName" comment:"商品名称"`
	Specification  string  `json:"specification" comment:"商品规格"`
	Unit           string  `json:"unit" comment:"单位"`
	Num            int64   `json:"num" comment:"数量"`
	Price          float64 `json:"price" comment:"单价"`
	TotalPrice     float64 `json:"totalPrice" comment:"总价"`
	PurchaseListId int64   `json:"purchaseListId" comment:"采购单"`
	GoodsId        int64   `json:"goodsId" comment:"商品"`
	SupplierId     int64   `json:"supplierId" comment:"供货商"`
	common.ControlBy
}

func (s *PurchaseListGoodsInsertReq) Generate(model *models.PurchaseListGoods) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GCode = s.GCode
	model.GName = s.GName
	model.Specification = s.Specification
	model.Unit = s.Unit
	model.Num = s.Num
	model.Price = s.Price
	model.TotalPrice = s.TotalPrice
	model.PurchaseListId = s.PurchaseListId
	model.GoodsId = s.GoodsId
	model.SupplierId = s.SupplierId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *PurchaseListGoodsInsertReq) GetId() interface{} {
	return s.Id
}

type PurchaseListGoodsUpdateReq struct {
	Id             int     `uri:"id" comment:"主键编码"` // 主键编码
	GCode          string  `json:"gCode" comment:"商品编码"`
	GName          string  `json:"gName" comment:"商品名称"`
	Specification  string  `json:"specification" comment:"商品规格"`
	Unit           string  `json:"unit" comment:"单位"`
	Num            int64   `json:"num" comment:"数量"`
	Price          float64 `json:"price" comment:"单价"`
	TotalPrice     float64 `json:"totalPrice" comment:"总价"`
	PurchaseListId int64   `json:"purchaseListId" comment:"采购单"`
	GoodsId        int64   `json:"goodsId" comment:"商品"`
	SupplierId     int64   `json:"supplierId" comment:"供货商"`
	common.ControlBy
}

func (s *PurchaseListGoodsUpdateReq) Generate(model *models.PurchaseListGoods) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.GCode = s.GCode
	model.GName = s.GName
	model.Specification = s.Specification
	model.Unit = s.Unit
	model.Num = s.Num
	model.Price = s.Price
	model.TotalPrice = s.TotalPrice
	model.PurchaseListId = s.PurchaseListId
	model.GoodsId = s.GoodsId
	model.SupplierId = s.SupplierId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *PurchaseListGoodsUpdateReq) GetId() interface{} {
	return s.Id
}

// PurchaseListGoodsGetReq 功能获取请求参数
type PurchaseListGoodsGetReq struct {
	Id int `uri:"id"`
}

func (s *PurchaseListGoodsGetReq) GetId() interface{} {
	return s.Id
}

// PurchaseListGoodsDeleteReq 功能删除请求参数
type PurchaseListGoodsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *PurchaseListGoodsDeleteReq) GetId() interface{} {
	return s.Ids
}
