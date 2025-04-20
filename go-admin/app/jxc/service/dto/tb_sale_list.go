package dto

import (
	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SaleListGetPageReq struct {
	dto.Pagination `search:"-"`
	SaleNumber     string `form:"saleNumber"  search:"type:contains;column:sale_number;table:tb_sale_list" comment:"销售单号"`
	State          string `form:"state"  search:"type:exact;column:state;table:tb_sale_list" comment:"状态"`
	CustomerId     int64  `form:"customerId"  search:"type:exact;column:customer_id;table:tb_sale_list" comment:"消费者ID"`
	ShipmentStatus string `form:"shipmentStatus"  search:"type:exact;column:shipment_status;table:tb_sale_list" comment:"出库状态(已出库/未出库)"`
	IsPurchased    string `form:"isPurchased"  search:"type:exact;column:is_purchased;comment:采购状态（0未采购，1已采购）"`
	SaleListOrder
}

type SaleListOrder struct {
	Id             string `form:"idOrder"  search:"type:order;column:id;table:tb_sale_list"`
	AmountPayable  string `form:"amountPayableOrder"  search:"type:order;column:amount_payable;table:tb_sale_list"`
	AmountPaid     string `form:"amountPaidOrder"  search:"type:order;column:amount_paid;table:tb_sale_list"`
	SaleDate       string `form:"saleDateOrder"  search:"type:order;column:sale_date;table:tb_sale_list"`
	SaleNumber     string `form:"saleNumberOrder"  search:"type:order;column:sale_number;table:tb_sale_list"`
	Remarks        string `form:"remarksOrder"  search:"type:order;column:remarks;table:tb_sale_list"`
	State          string `form:"stateOrder"  search:"type:order;column:state;table:tb_sale_list"`
	ShipmentStatus string `form:"shipmentStatus"  search:"type:order;column:shipment_status;table:tb_sale_list"`
	CustomerId     string `form:"customerIdOrder"  search:"type:order;column:customer_id;table:tb_sale_list"`
	CreatedAt      string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_sale_list"`
	UpdatedAt      string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_sale_list"`
	DeletedAt      string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_sale_list"`
	CreateBy       string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_sale_list"`
	UpdateBy       string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_sale_list"`
	IsPurchased    string `form:"isPurchased"  search:"type:order;column:is_purchased;table:tb_sale_list"`
}

func (m *SaleListGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SaleListInsertReq struct {
	Id            int                      `json:"-" comment:"主键编码"` // 主键编码
	AmountPayable float64                  `json:"amountPayable" comment:"应付金额"`
	AmountPaid    float64                  `json:"amountPaid" comment:"已付金额"`
	SaleDate      string                   `json:"saleDate" comment:"销售日期"`
	SaleNumber    string                   `json:"saleNumber" comment:"销售单号"`
	Remarks       string                   `json:"remarks" comment:"备注"`
	State         string                   `json:"state" comment:"状态"`
	CustomerId    int64                    `json:"customerId" comment:"消费者ID"`
	Goods         []SaleListGoodsInsertReq `json:"goods" comment:"商品列表"`
	common.ControlBy
}

func (s *SaleListInsertReq) Generate(model *models.SaleList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AmountPayable = s.AmountPayable
	model.AmountPaid = s.AmountPaid
	model.SaleDate = s.SaleDate
	model.SaleNumber = s.SaleNumber
	model.Remarks = s.Remarks
	model.State = s.State
	model.CustomerId = s.CustomerId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.IsPurchased = "0"
}
func (s *SaleListInsertReq) GenerateGoods(saleListId int64) []models.SaleListGoods {
	goodsDtos := s.Goods
	var slgoodsList []models.SaleListGoods
	for _, g := range goodsDtos {
		slGoods := models.SaleListGoods{}
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
		slGoods.SaleListId = saleListId
		slGoods.GoodsId = g.GoodsId
		slGoods.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
		slgoodsList = append(slgoodsList, slGoods)
	}
	return slgoodsList
}

func (s *SaleListInsertReq) GetId() interface{} {
	return s.Id
}

type SaleListUpdateReq struct {
	Id            int     `uri:"id" comment:"主键编码"` // 主键编码
	AmountPayable float64 `json:"amountPayable" comment:"应付金额"`
	AmountPaid    float64 `json:"amountPaid" comment:"已付金额"`
	SaleDate      string  `json:"saleDate" comment:"销售日期"`
	SaleNumber    string  `json:"saleNumber" comment:"销售单号"`
	Remarks       string  `json:"remarks" comment:"备注"`
	State         string  `json:"state" comment:"状态"`
	CustomerId    int64   `json:"customerId" comment:"消费者ID"`
	common.ControlBy
}

func (s *SaleListUpdateReq) Generate(model *models.SaleList) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.AmountPayable = s.AmountPayable
	model.AmountPaid = s.AmountPaid
	model.SaleDate = s.SaleDate
	model.SaleNumber = s.SaleNumber
	model.Remarks = s.Remarks
	model.State = s.State
	model.CustomerId = s.CustomerId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SaleListUpdateReq) GetId() interface{} {
	return s.Id
}

// SaleListGetReq 功能获取请求参数
type SaleListGetReq struct {
	Id int `uri:"id"`
}

func (s *SaleListGetReq) GetId() interface{} {
	return s.Id
}

// SaleListDeleteReq 功能删除请求参数
type SaleListDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SaleListDeleteReq) GetId() interface{} {
	return s.Ids
}
