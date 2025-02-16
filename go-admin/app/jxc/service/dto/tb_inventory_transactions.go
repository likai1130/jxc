package dto

import (
	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type InventoryTransactionsGetPageReq struct {
	dto.Pagination  `search:"-"`
	TransactionType string `form:"transactionType"  search:"type:exact;column:transaction_type;table:tb_inventory_transactions" comment:"操作类型"`
	GCode           string `form:"gCode"  search:"type:contains;column:g_code;table:tb_inventory_transactions" comment:"商品编码"`
	GName           string `form:"gName"  search:"type:contains;column:g_name;table:tb_inventory_transactions" comment:"商品名称"`
	OrderNumber     string `form:"orderNumber"  search:"type:exact;column:order_number;table:tb_inventory_transactions" comment:"订单号"`
	GoodsId         int64  `form:"goodsId"  search:"type:exact;column:goods_id;table:tb_inventory_transactions" comment:"商品ID"`
	OrderId         int64  `form:"orderId"  search:"type:exact;column:order_id;table:tb_inventory_transactions" comment:"订单ID"`
	InventoryTransactionsOrder
}

type InventoryTransactionsOrder struct {
	Id              string `form:"idOrder"  search:"type:order;column:id;table:tb_inventory_transactions"`
	TransactionType string `form:"transactionTypeOrder"  search:"type:order;column:transaction_type;table:tb_inventory_transactions"`
	GCode           string `form:"gCodeOrder"  search:"type:order;column:g_code;table:tb_inventory_transactions"`
	GName           string `form:"gNameOrder"  search:"type:order;column:g_name;table:tb_inventory_transactions"`
	Num             string `form:"numOrder"  search:"type:order;column:num;table:tb_inventory_transactions"`
	OrderNumber     string `form:"orderNumberOrder"  search:"type:order;column:order_number;table:tb_inventory_transactions"`
	GoodsId         string `form:"goodsIdOrder"  search:"type:order;column:goods_id;table:tb_inventory_transactions"`
	OrderId         string `form:"orderIdOrder"  search:"type:order;column:order_id;table:tb_inventory_transactions"`
	Remarks         string `form:"remarksOrder"  search:"type:order;column:remarks;table:tb_inventory_transactions"`
	CreatedAt       string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_inventory_transactions"`
	UpdatedAt       string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_inventory_transactions"`
	DeletedAt       string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_inventory_transactions"`
	CreateBy        string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_inventory_transactions"`
	UpdateBy        string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_inventory_transactions"`
}

func (m *InventoryTransactionsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type InventoryTransactionsInsertReq struct {
	Id              int     `json:"-" comment:"主键编码"` // 主键编码
	TransactionType string  `json:"transactionType" comment:"操作类型"`
	OrderNumber     string  `json:"orderNumber" comment:"订单号"`
	OrderId         int64   `json:"orderId" comment:"订单ID"`
	Remarks         string  `json:"remarks" comment:"备注"`
	GoodsList       []Goods `json:"goods" comment:"入库商品列表"`
	common.ControlBy
}

type Goods struct {
	GCode               string  `json:"gCode" comment:"商品编码"`
	GName               string  `json:"gName" comment:"商品名称"`
	Specification       string  ` json:"specification" comment:"商品规格"`
	Num                 int64   `json:"num" comment:"数量"`
	LastPurchasingPrice float64 `json:"price" comment:"当前采购价"`
	GoodsId             int64   `json:"goodsId" comment:"商品ID"`
}

func (s *InventoryTransactionsInsertReq) Generate() []models.InventoryTransactions {
	var inventoryTransactionsList []models.InventoryTransactions
	if len(s.GoodsList) > 0 {
		for _, gl := range s.GoodsList {
			transactions := models.InventoryTransactions{
				TransactionType: s.TransactionType,
				GCode:           gl.GCode,
				GName:           gl.GCode,
				Specification:   gl.Specification,
				Num:             gl.Num,
				OrderNumber:     s.OrderNumber,
				GoodsId:         gl.GoodsId,
				OrderId:         s.OrderId,
				Remarks:         s.Remarks,
			}
			transactions.CreateBy = s.CreateBy
			inventoryTransactionsList = append(inventoryTransactionsList, transactions)
		}
		return inventoryTransactionsList
	}
	return nil
}

func (s *InventoryTransactionsInsertReq) GetId() interface{} {
	return s.Id
}

type InventoryTransactionsUpdateReq struct {
	Id              int    `uri:"id" comment:"主键编码"` // 主键编码
	TransactionType string `json:"transactionType" comment:"操作类型"`
	GCode           string `json:"gCode" comment:"商品编码"`
	GName           string `json:"gName" comment:"商品名称"`
	Num             int64  `json:"num" comment:"数量"`
	OrderNumber     string `json:"orderNumber" comment:"订单号"`
	GoodsId         int64  `json:"goodsId" comment:"商品ID"`
	OrderId         int64  `json:"orderId" comment:"订单ID"`
	Remarks         string `json:"remarks" comment:"备注"`
	common.ControlBy
}

func (s *InventoryTransactionsUpdateReq) Generate(model *models.InventoryTransactions) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.TransactionType = s.TransactionType
	model.GCode = s.GCode
	model.GName = s.GName
	model.Num = s.Num
	model.OrderNumber = s.OrderNumber
	model.GoodsId = s.GoodsId
	model.OrderId = s.OrderId
	model.Remarks = s.Remarks
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *InventoryTransactionsUpdateReq) GetId() interface{} {
	return s.Id
}

// InventoryTransactionsGetReq 功能获取请求参数
type InventoryTransactionsGetReq struct {
	Id int `uri:"id"`
}

func (s *InventoryTransactionsGetReq) GetId() interface{} {
	return s.Id
}

// InventoryTransactionsDeleteReq 功能删除请求参数
type InventoryTransactionsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *InventoryTransactionsDeleteReq) GetId() interface{} {
	return s.Ids
}
