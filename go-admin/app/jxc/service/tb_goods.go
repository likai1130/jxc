package service

import (
	"errors"
	"fmt"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/jxc/models"
	"go-admin/app/jxc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type Goods struct {
	service.Service
}

// GetPage 获取Goods列表
func (e *Goods) GetPage(c *dto.GoodsGetPageReq, p *actions.DataPermission, list *[]models.Goods, count *int64) error {
	var err error
	var data models.Goods

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("GoodsService GetPage error:%s \r\n", err)
		return err
	}
	// 计算销售总数
	if len(*list) > 0 {
		for i, goods := range *list {
			var totalSales int64
			err = e.Orm.Model(models.SaleListGoods{}).
				Select("SUM(num)").
				Where("goods_id = ?", goods.Id).
				Scan(&totalSales).Error
			if err != nil {
				return err
			}
			m := *list
			m[i].SaleNum = totalSales
		}
	}
	return nil
}

// Get 获取Goods对象
func (e *Goods) Get(d *dto.GoodsGetReq, p *actions.DataPermission, model *models.Goods) error {
	var data models.Goods

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetGoods error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Goods对象
func (e *Goods) Insert(c *dto.GoodsInsertReq) error {
	var err error
	var data models.Goods
	c.Generate(&data)

	var productCount int64
	if err = e.Orm.Model(&models.Goods{}).Count(&productCount).Error; err != nil {
		return err
	}

	// 计算下一个产品 ID
	nextProductID := productCount + 1
	// 格式化产品 ID 为 5 位
	formattedProductID := fmt.Sprintf("%05d", nextProductID)
	data.GCode = formattedProductID

	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("GoodsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Goods对象
func (e *Goods) Update(c *dto.GoodsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Goods{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("GoodsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Goods
func (e *Goods) Remove(d *dto.GoodsDeleteReq, p *actions.DataPermission) error {
	var data models.Goods

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveGoods error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
