package dto

import (
	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type PurchaseListGetPageReq struct {
	dto.Pagination `search:"-"`
	PurchaseNumber string `form:"purchaseNumber"  search:"type:exact;column:purchase_number;table:tb_purchase_list" comment:"采购单号"`
	State          string `form:"state"  search:"type:exact;column:state;table:tb_purchase_list" comment:"是否付款"`
	StorageStatus  string `form:"storageStatus"  search:"type:exact;column:storage_status;table:tb_purchase_list" comment:"是否入库"`
	PurchaseListOrder
}

type PurchaseListOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:tb_purchase_list"`
	AmountPayable  string `form:"amountPayableOrder"  search:"type:order;column:amount_payable;table:tb_purchase_list"`
	AmountPaid     string `form:"amountPaidOrder"  search:"type:order;column:amount_paid;table:tb_purchase_list"`
	PurchaseDate   string `form:"purchaseDateOrder"  search:"type:order;column:purchase_date;table:tb_purchase_list"`
	PurchaseNumber string `form:"purchaseNumberOrder"  search:"type:order;column:purchase_number;table:tb_purchase_list"`
	Remarks        string `form:"remarksOrder"  search:"type:order;column:remarks;table:tb_purchase_list"`
	State          string `form:"stateOrder"  search:"type:order;column:state;table:tb_purchase_list"`
	StorageStatus  string `form:"storageStatusOrder"  search:"type:order;column:storage_status;table:tb_purchase_list"`
	SupplierId     string `form:"supplierIdOrder"  search:"type:order;column:supplier_id;table:tb_purchase_list"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_purchase_list"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_purchase_list"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_purchase_list"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_purchase_list"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_purchase_list"`
}

func (m *PurchaseListGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type PurchaseListInsertReq struct {
	Id             int                          `json:"-" comment:"主键编码"` // 主键编码
	AmountPayable  float64                      `json:"amountPayable" comment:"应付金额"`
	AmountPaid     float64                      `json:"amountPaid" comment:"已付金额"`
	PurchaseDate   string                       `json:"purchaseDate" comment:"采购日期"`
	PurchaseNumber string                       `json:"purchaseNumber" comment:"采购单号"`
	Remarks        string                       `json:"remarks" comment:"备注"`
	State          string                       `json:"state" comment:"是否付款"`
	StorageStatus  string                       `json:"storageStatus" comment:"是否入库"`
	SupplierId     int64                        `json:"supplierId" comment:"供货商"`
	Goods          []PurchaseListGoodsInsertReq `json:"goods" comment:"商品列表"`
	common.ControlBy
}

func (s *PurchaseListInsertReq) Generate(model *models.PurchaseList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AmountPayable = s.AmountPayable
	model.AmountPaid = s.AmountPaid
	model.PurchaseDate = s.PurchaseDate
	model.PurchaseNumber = s.PurchaseNumber
	model.Remarks = s.Remarks
	model.State = s.State
	model.StorageStatus = s.StorageStatus
	model.SupplierId = s.SupplierId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *PurchaseListInsertReq) GenerateGoods(purchaseListId int64) []models.PurchaseListGoods {
	goodsDtos := s.Goods
	var slgoodsList []models.PurchaseListGoods
	for _, g := range goodsDtos {
		slGoods := models.PurchaseListGoods{}
		if g.Id == 0 {
			slGoods.Model = common.Model{Id: g.Id}
		}

		slGoods.GCode = g.GCode
		slGoods.GName = g.GName
		slGoods.Specification = g.Specification
		slGoods.Unit = g.Unit
		slGoods.Num = g.Num
		slGoods.Price = g.Price
		slGoods.TotalPrice = g.TotalPrice
		slGoods.PurchaseListId = purchaseListId
		slGoods.SupplierId = g.SupplierId
		slGoods.GoodsId = g.GoodsId
		slGoods.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
		slgoodsList = append(slgoodsList, slGoods)
	}
	return slgoodsList
}

func (s *PurchaseListInsertReq) GetId() interface{} {
	return s.Id
}

type PurchaseListUpdateReq struct {
	Id             int     `uri:"id" comment:"主键编码"` // 主键编码
	AmountPayable  float64 `json:"amountPayable" comment:"应付金额"`
	AmountPaid     float64 `json:"amountPaid" comment:"已付金额"`
	PurchaseDate   string  `json:"purchaseDate" comment:"采购日期"`
	PurchaseNumber string  `json:"purchaseNumber" comment:"采购单号"`
	Remarks        string  `json:"remarks" comment:"备注"`
	State          string  `json:"state" comment:"是否付款"`
	StorageStatus  string  `json:"storageStatus" comment:"是否入库"`
	SupplierId     int64   `json:"supplierId" comment:"供货商"`
	common.ControlBy
}

func (s *PurchaseListUpdateReq) Generate(model *models.PurchaseList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AmountPayable = s.AmountPayable
	model.AmountPaid = s.AmountPaid
	model.PurchaseDate = s.PurchaseDate
	model.PurchaseNumber = s.PurchaseNumber
	model.Remarks = s.Remarks
	model.State = s.State
	model.StorageStatus = s.StorageStatus
	model.SupplierId = s.SupplierId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *PurchaseListUpdateReq) GetId() interface{} {
	return s.Id
}

// PurchaseListGetReq 功能获取请求参数
type PurchaseListGetReq struct {
	Id int `uri:"id"`
}

func (s *PurchaseListGetReq) GetId() interface{} {
	return s.Id
}

// PurchaseListDeleteReq 功能删除请求参数
type PurchaseListDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *PurchaseListDeleteReq) GetId() interface{} {
	return s.Ids
}
