package dto

import (
	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ConsumerGetPageReq struct {
	dto.Pagination `search:"-"`
	CustomerName   string `form:"customerName"  search:"type:contains;column:customer_name;table:tb_consumer" comment:"客户名称"`
	CustomerType   string `form:"customerType"  search:"type:exact;column:customer_type;table:tb_consumer" comment:"客户类型"`
	ConsumerOrder
}

type ConsumerOrder struct {
	Id              string `form:"idOrder"  search:"type:order;column:id;table:tb_consumer"`
	CustomerName    string `form:"customerNameOrder"  search:"type:order;column:customer_name;table:tb_consumer"`
	CustomerType    string `form:"customerTypeOrder"  search:"type:order;column:customer_type;table:tb_consumer"`
	ContactName     string `form:"contactNameOrder"  search:"type:order;column:contact_name;table:tb_consumer"`
	ContactPhone    string `form:"contactPhoneOrder"  search:"type:order;column:contact_phone;table:tb_consumer"`
	Email           string `form:"emailOrder"  search:"type:order;column:email;table:tb_consumer"`
	CompanyAddress  string `form:"companyAddressOrder"  search:"type:order;column:company_address;table:tb_consumer"`
	DeliveryAddress string `form:"deliveryAddressOrder"  search:"type:order;column:delivery_address;table:tb_consumer"`
	Notes           string `form:"notesOrder"  search:"type:order;column:notes;table:tb_consumer"`
	CreatedAt       string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_consumer"`
	UpdatedAt       string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_consumer"`
	DeletedAt       string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_consumer"`
	CreateBy        string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_consumer"`
	UpdateBy        string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_consumer"`
}

func (m *ConsumerGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ConsumerInsertReq struct {
	Id              int    `json:"-" comment:"客户ID，自动递增"` // 客户ID，自动递增
	CustomerName    string `json:"customerName" comment:"客户名称"`
	CustomerType    string `json:"customerType" comment:"客户类型"`
	ContactName     string `json:"contactName" comment:"联系人姓名"`
	ContactPhone    string `json:"contactPhone" comment:"联系电话"`
	Email           string `json:"email" comment:"电子邮箱"`
	CompanyAddress  string `json:"companyAddress" comment:"公司地址"`
	DeliveryAddress string `json:"deliveryAddress" comment:"送货地址"`
	Notes           string `json:"notes" comment:"备注信息"`
	common.ControlBy
}

func (s *ConsumerInsertReq) Generate(model *models.Consumer) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CustomerName = s.CustomerName
	model.CustomerType = s.CustomerType
	model.ContactName = s.ContactName
	model.ContactPhone = s.ContactPhone
	model.Email = s.Email
	model.CompanyAddress = s.CompanyAddress
	model.DeliveryAddress = s.DeliveryAddress
	model.Notes = s.Notes
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ConsumerInsertReq) GetId() interface{} {
	return s.Id
}

type ConsumerUpdateReq struct {
	Id              int    `uri:"id" comment:"客户ID，自动递增"` // 客户ID，自动递增
	CustomerName    string `json:"customerName" comment:"客户名称"`
	CustomerType    string `json:"customerType" comment:"客户类型"`
	ContactName     string `json:"contactName" comment:"联系人姓名"`
	ContactPhone    string `json:"contactPhone" comment:"联系电话"`
	Email           string `json:"email" comment:"电子邮箱"`
	CompanyAddress  string `json:"companyAddress" comment:"公司地址"`
	DeliveryAddress string `json:"deliveryAddress" comment:"送货地址"`
	Notes           string `json:"notes" comment:"备注信息"`
	common.ControlBy
}

func (s *ConsumerUpdateReq) Generate(model *models.Consumer) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.CustomerName = s.CustomerName
	model.CustomerType = s.CustomerType
	model.ContactName = s.ContactName
	model.ContactPhone = s.ContactPhone
	model.Email = s.Email
	model.CompanyAddress = s.CompanyAddress
	model.DeliveryAddress = s.DeliveryAddress
	model.Notes = s.Notes
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ConsumerUpdateReq) GetId() interface{} {
	return s.Id
}

// ConsumerGetReq 功能获取请求参数
type ConsumerGetReq struct {
	Id int `uri:"id"`
}

func (s *ConsumerGetReq) GetId() interface{} {
	return s.Id
}

// ConsumerDeleteReq 功能删除请求参数
type ConsumerDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ConsumerDeleteReq) GetId() interface{} {
	return s.Ids
}
