package models

type Supplier struct {
	Model
	SupplierName string `gorm:"not null;type:varchar(255);comment:供应商名称" json:"supplierName"` // 供应商名称
	ContactName  string `gorm:"type:varchar(255);comment:联系人姓名" json:"contactName"`           // 联系人姓名
	PhoneNumber  string `gorm:"type:varchar(50);comment:联系人电话" json:"phoneNumber"`            // 联系人电话
	Address      string `gorm:"type:varchar(255);comment:供应商地址" json:"address"`               // 供应商地址
	Remarks      string `gorm:"type:text;comment:备注" json:"remarks"`                          // 备注
	ModelTime
	ControlBy
}

func (Supplier) TableName() string {
	return "tb_supplier"
}
