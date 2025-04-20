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

type PurchaseList struct {
	service.Service
}

// GetPage 获取PurchaseList列表
func (e *PurchaseList) GetPage(c *dto.PurchaseListGetPageReq, p *actions.DataPermission, list *[]models.PurchaseList, count *int64) error {
	var err error
	var data models.PurchaseList

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("PurchaseListService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取PurchaseList对象
func (e *PurchaseList) Get(d *dto.PurchaseListGetReq, p *actions.DataPermission, model *models.PurchaseList) error {
	var data models.PurchaseList

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetPurchaseList error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建PurchaseList对象
func (e *PurchaseList) Insert(c *dto.PurchaseListInsertReq) error {
	var err error
	var data models.PurchaseList

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

	// 1. 生成新的采购单号
	// 获取今天的开始和结束时间
	now := time.Now()
	startOfDay := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	formattedStartTime := startOfDay.Format("2006-01-02 15:04:05")
	endOfDay := startOfDay.Add(24 * time.Hour)
	formattedEndTime := endOfDay.Format("2006-01-02 15:04:05")

	var count int64
	err = tx.Model(&models.PurchaseList{}).Where("created_at >= ? AND created_at < ?", formattedStartTime, formattedEndTime).Count(&count).Error
	if err != nil {
		return err
	}
	// 生成订单号
	saleNumber := utils.GeneratePurchaseOrderNumber(count)
	data.PurchaseNumber = saleNumber
	data.StorageStatus = "0"
	// 2. 创建订单
	err = tx.Create(&data).Error
	if err != nil {
		e.Log.Errorf("PurchaseListService Insert error:%s", err)
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
	// 4. 修改销售单的采购状态
	if err = tx.Model(&models.SaleList{}).Where("sale_number = ?", c.SelectedSaleNumberValue).Updates(map[string]interface{}{"is_purchased": "1"}).Error; err != nil {
		e.Log.Errorf("update SaleList is_purchased error:%s", err)
		return err
	}
	return nil
}

// Update 修改PurchaseList对象
func (e *PurchaseList) Update(c *dto.PurchaseListUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.PurchaseList{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("PurchaseListService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除PurchaseList
func (e *PurchaseList) Remove(d *dto.PurchaseListDeleteReq, p *actions.DataPermission) error {
	var data models.PurchaseList

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemovePurchaseList error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
