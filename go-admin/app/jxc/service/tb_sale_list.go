package service

import (
	"errors"
	"go-admin/app/jxc/utils"
	"time"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/jxc/models"
	"go-admin/app/jxc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type SaleList struct {
	service.Service
}

// GetPage 获取SaleList列表
func (e *SaleList) GetPage(c *dto.SaleListGetPageReq, p *actions.DataPermission, list *[]models.SaleList, count *int64) error {
	var err error
	var data models.SaleList

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("SaleListService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取SaleList对象
func (e *SaleList) Get(d *dto.SaleListGetReq, p *actions.DataPermission, model *models.SaleList) error {
	var data models.SaleList

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetSaleList error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建SaleList对象
func (e *SaleList) Insert(c *dto.SaleListInsertReq) error {
	var err error
	var data models.SaleList
	c.Generate(&data)
	// 创建事务
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// 1. 生成新的销售单号
	// 获取今天的开始和结束时间
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	formattedStartTime := startOfDay.Format("2006-01-02 15:04:05")
	endOfDay := startOfDay.Add(24 * time.Hour)
	formattedEndTime := endOfDay.Format("2006-01-02 15:04:05")

	var count int64
	err = tx.Model(&models.SaleList{}).Where("created_at >= ? AND created_at < ?", formattedStartTime, formattedEndTime).Count(&count).Error
	if err != nil {
		return err
	}
	// 生成订单号
	saleNumber := utils.GenerateSalesOrderNumber(count)
	data.SaleNumber = saleNumber
	data.ShipmentStatus = "0"
	// 2. 创建订单
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("SaleListService Insert error:%s \r\n", err)
		return err
	}
	// 3. 保存商品明细
	generateGoods := c.GenerateGoods(int64(data.Id))
	if len(generateGoods) > 0 {
		err = tx.CreateInBatches(generateGoods, len(generateGoods)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

// Update 修改SaleList对象
func (e *SaleList) Update(c *dto.SaleListUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.SaleList{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("SaleListService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除SaleList
func (e *SaleList) Remove(d *dto.SaleListDeleteReq, p *actions.DataPermission) error {
	var data models.SaleList

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveSaleList error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
