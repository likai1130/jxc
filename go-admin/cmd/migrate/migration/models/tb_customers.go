package models

type Customer struct {
	Model
	CustomerName    string `gorm:"column:customer_name;type:varchar(255);not null;comment:客户名称" json:"customer_name"`           // 客户名称
	CustomerType    string `gorm:"column:customer_type;type:varchar(100);default:null;comment:客户类型" json:"customer_type"`       // 客户类型
	ContactName     string `gorm:"column:contact_name;type:varchar(100);default:null;comment:联系人姓名" json:"contact_name"`        // 联系人姓名
	ContactPhone    string `gorm:"column:contact_phone;type:varchar(20);default:null;comment:联系电话" json:"contact_phone"`        // 联系电话
	Email           string `gorm:"column:email;type:varchar(100);default:null;comment:电子邮箱" json:"email"`                       // 电子邮箱
	CompanyAddress  string `gorm:"column:company_address;type:varchar(255);default:null;comment:公司地址" json:"company_address"`   // 公司地址
	DeliveryAddress string `gorm:"column:delivery_address;type:varchar(255);default:null;comment:送货地址" json:"delivery_address"` // 送货地址
	Notes           string `gorm:"column:notes;type:varchar(255);default:null;comment:备注信息" json:"notes"`                       // 备注信息
	ModelTime
	ControlBy
}

func (Customer) TableName() string {
	return "tb_consumer"
}
