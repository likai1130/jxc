package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/jxc/models"
	"go-admin/app/jxc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Supplier struct {
	service.Service
}

// GetPage 获取Supplier列表
func (e *Supplier) GetPage(c *dto.SupplierGetPageReq, p *actions.DataPermission, list *[]models.Supplier, count *int64) error {
	var err error
	var data models.Supplier

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SupplierService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Supplier对象
func (e *Supplier) Get(d *dto.SupplierGetReq, p *actions.DataPermission, model *models.Supplier) error {
	var data models.Supplier

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSupplier error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Supplier对象
func (e *Supplier) Insert(c *dto.SupplierInsertReq) error {
    var err error
    var data models.Supplier
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SupplierService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Supplier对象
func (e *Supplier) Update(c *dto.SupplierUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Supplier{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SupplierService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Supplier
func (e *Supplier) Remove(d *dto.SupplierDeleteReq, p *actions.DataPermission) error {
	var data models.Supplier

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSupplier error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
