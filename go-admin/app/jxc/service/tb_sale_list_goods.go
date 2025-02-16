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

type SaleListGoods struct {
	service.Service
}

// GetPage 获取SaleListGoods列表
func (e *SaleListGoods) GetPage(c *dto.SaleListGoodsGetPageReq, p *actions.DataPermission, list *[]models.SaleListGoods, count *int64) error {
	var err error
	var data models.SaleListGoods

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SaleListGoodsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SaleListGoods对象
func (e *SaleListGoods) Get(d *dto.SaleListGoodsGetReq, p *actions.DataPermission, model *models.SaleListGoods) error {
	var data models.SaleListGoods

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSaleListGoods error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SaleListGoods对象
func (e *SaleListGoods) Insert(c *dto.SaleListGoodsInsertReq) error {
    var err error
    var data models.SaleListGoods
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SaleListGoodsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改SaleListGoods对象
func (e *SaleListGoods) Update(c *dto.SaleListGoodsUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.SaleListGoods{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("SaleListGoodsService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除SaleListGoods
func (e *SaleListGoods) Remove(d *dto.SaleListGoodsDeleteReq, p *actions.DataPermission) error {
	var data models.SaleListGoods

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveSaleListGoods error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
