package models

import (

	"go-admin/common/models"

)

type Supplier struct {
    models.Model
    
    SupplierName string `json:"supplierName" gorm:"type:varchar(255);comment:供应商名称"` 
    ContactName string `json:"contactName" gorm:"type:varchar(255);comment:联系人名称"` 
    PhoneNumber string `json:"phoneNumber" gorm:"type:varchar(50);comment:联系人电话"` 
    Address string `json:"address" gorm:"type:varchar(255);comment:供应商地址"` 
    Remarks string `json:"remarks" gorm:"type:text;comment:备注"` 
    models.ModelTime
    models.ControlBy
}

func (Supplier) TableName() string {
    return "tb_supplier"
}

func (e *Supplier) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Supplier) GetId() interface{} {
	return e.Id
}