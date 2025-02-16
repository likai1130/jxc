package dto

import (

	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type SupplierGetPageReq struct {
	dto.Pagination     `search:"-"`
    SupplierName string `form:"supplierName"  search:"type:contains;column:supplier_name;table:tb_supplier" comment:"供应商名称"`
    ContactName string `form:"contactName"  search:"type:contains;column:contact_name;table:tb_supplier" comment:"联系人名称"`
    SupplierOrder
}

type SupplierOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_supplier"`
    SupplierName string `form:"supplierNameOrder"  search:"type:order;column:supplier_name;table:tb_supplier"`
    ContactName string `form:"contactNameOrder"  search:"type:order;column:contact_name;table:tb_supplier"`
    PhoneNumber string `form:"phoneNumberOrder"  search:"type:order;column:phone_number;table:tb_supplier"`
    Address string `form:"addressOrder"  search:"type:order;column:address;table:tb_supplier"`
    Remarks string `form:"remarksOrder"  search:"type:order;column:remarks;table:tb_supplier"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_supplier"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_supplier"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_supplier"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_supplier"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_supplier"`
    
}

func (m *SupplierGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type SupplierInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    SupplierName string `json:"supplierName" comment:"供应商名称"`
    ContactName string `json:"contactName" comment:"联系人名称"`
    PhoneNumber string `json:"phoneNumber" comment:"联系人电话"`
    Address string `json:"address" comment:"供应商地址"`
    Remarks string `json:"remarks" comment:"备注"`
    common.ControlBy
}

func (s *SupplierInsertReq) Generate(model *models.Supplier)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.SupplierName = s.SupplierName
    model.ContactName = s.ContactName
    model.PhoneNumber = s.PhoneNumber
    model.Address = s.Address
    model.Remarks = s.Remarks
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *SupplierInsertReq) GetId() interface{} {
	return s.Id
}

type SupplierUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    SupplierName string `json:"supplierName" comment:"供应商名称"`
    ContactName string `json:"contactName" comment:"联系人名称"`
    PhoneNumber string `json:"phoneNumber" comment:"联系人电话"`
    Address string `json:"address" comment:"供应商地址"`
    Remarks string `json:"remarks" comment:"备注"`
    common.ControlBy
}

func (s *SupplierUpdateReq) Generate(model *models.Supplier)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.SupplierName = s.SupplierName
    model.ContactName = s.ContactName
    model.PhoneNumber = s.PhoneNumber
    model.Address = s.Address
    model.Remarks = s.Remarks
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *SupplierUpdateReq) GetId() interface{} {
	return s.Id
}

// SupplierGetReq 功能获取请求参数
type SupplierGetReq struct {
     Id int `uri:"id"`
}
func (s *SupplierGetReq) GetId() interface{} {
	return s.Id
}

// SupplierDeleteReq 功能删除请求参数
type SupplierDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *SupplierDeleteReq) GetId() interface{} {
	return s.Ids
}
