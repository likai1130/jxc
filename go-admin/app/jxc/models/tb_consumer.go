package models

import (
	"go-admin/common/models"
)

type Consumer struct {
	models.Model
	CustomerName    string `json:"customerName" gorm:"type:varchar(255);comment:客户名称"`
	CustomerType    string `json:"customerType" gorm:"type:varchar(100);comment:客户类型"`
	ContactName     string `json:"contactName" gorm:"type:varchar(100);comment:联系人姓名"`
	ContactPhone    string `json:"contactPhone" gorm:"type:varchar(20);comment:联系电话"`
	Email           string `json:"email" gorm:"type:varchar(100);comment:电子邮箱"`
	CompanyAddress  string `json:"companyAddress" gorm:"type:varchar(255);comment:公司地址"`
	DeliveryAddress string `json:"deliveryAddress" gorm:"type:varchar(255);comment:送货地址"`
	Notes           string `json:"notes" gorm:"type:varchar(255);comment:备注信息"`
	models.ModelTime
	models.ControlBy
}

func (Consumer) TableName() string {
	return "tb_consumer"
}

func (e *Consumer) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Consumer) GetId() interface{} {
	return e.Id
}
