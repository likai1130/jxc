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

type PurchaseListGoods struct {
	service.Service
}

// GetPage 获取PurchaseListGoods列表
func (e *PurchaseListGoods) GetPage(c *dto.PurchaseListGoodsGetPageReq, p *actions.DataPermission, list *[]models.PurchaseListGoods, count *int64) error {
	var err error
	var data models.PurchaseListGoods

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("PurchaseListGoodsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取PurchaseListGoods对象
func (e *PurchaseListGoods) Get(d *dto.PurchaseListGoodsGetReq, p *actions.DataPermission, model *models.PurchaseListGoods) error {
	var data models.PurchaseListGoods

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetPurchaseListGoods error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建PurchaseListGoods对象
func (e *PurchaseListGoods) Insert(c *dto.PurchaseListGoodsInsertReq) error {
    var err error
    var data models.PurchaseListGoods
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("PurchaseListGoodsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改PurchaseListGoods对象
func (e *PurchaseListGoods) Update(c *dto.PurchaseListGoodsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.PurchaseListGoods{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("PurchaseListGoodsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除PurchaseListGoods
func (e *PurchaseListGoods) Remove(d *dto.PurchaseListGoodsDeleteReq, p *actions.DataPermission) error {
	var data models.PurchaseListGoods

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemovePurchaseListGoods error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
