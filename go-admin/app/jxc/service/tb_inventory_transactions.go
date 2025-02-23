package service

import (
	"errors"
	"fmt"
	"go-admin/app/jxc/service/consts"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/jxc/models"
	"go-admin/app/jxc/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type InventoryTransactions struct {
	service.Service
}

// GetPage 获取InventoryTransactions列表
func (e *InventoryTransactions) GetPage(c *dto.InventoryTransactionsGetPageReq, p *actions.DataPermission, list *[]models.InventoryTransactions, count *int64) error {
	var err error
	var data models.InventoryTransactions

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("InventoryTransactionsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取InventoryTransactions对象
func (e *InventoryTransactions) Get(d *dto.InventoryTransactionsGetReq, p *actions.DataPermission, model *models.InventoryTransactions) error {
	var data models.InventoryTransactions

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetInventoryTransactions error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建InventoryTransactions对象
func (e *InventoryTransactions) Insert(c *dto.InventoryTransactionsInsertReq) error {
	var err error
	generateModelInventoryTransactions := c.Generate()
	tx := e.Orm.Begin()
	defer func() {
		if err != nil {
			tx.Callback()
		} else {
			tx.Commit()
		}
	}()
	// 1. 判断库存类型，如果是0入库，先修改商品表的均价，再增加数量，最后插入库存流水表，修改采购单状态
	if c.TransactionType == consts.Inventory {
		// 计算商品表的均价
		for _, gmit := range c.GoodsList {
			goodsData := models.Goods{}
			err = tx.First(&goodsData, gmit.GoodsId).Error
			if err != nil {
				return err
			}
			goodsData.PurchasingPrice = (goodsData.PurchasingPrice*float64(goodsData.StockQuantity) + gmit.LastPurchasingPrice*float64(gmit.Num)) / float64(goodsData.StockQuantity+gmit.Num)
			goodsData.LastPurchasingPrice = gmit.LastPurchasingPrice
			goodsData.StockQuantity += gmit.Num
			err = tx.Save(goodsData).Error
			if err != nil {
				return err
			}
		}
		err = tx.CreateInBatches(generateModelInventoryTransactions, len(generateModelInventoryTransactions)).Error
		if err != nil {
			return err
		}

		//  修改采购单状态
		err = tx.Model(&models.PurchaseList{}).Where("id = ?", c.OrderId).Update("storage_status", "1").Error
		if err != nil {
			return err
		}
		return nil
	}

	// 2. 判断库存类型，如果是1出库，减少数量，最后插入库存流水表，修改销售单状态
	if c.TransactionType == consts.Outbound {
		// 计算商品表的均价
		for _, gmit := range c.GoodsList {
			goodsData := models.Goods{}
			err = tx.First(&goodsData, gmit.GoodsId).Error
			if err != nil {
				return err
			}
			if goodsData.StockQuantity < gmit.Num {
				return errors.New(fmt.Sprintf("商品: %s 规格: %s 库存不足，请备货后再出库", gmit.GName, gmit.Specification))
			}
			goodsData.StockQuantity -= gmit.Num
			err = tx.Save(goodsData).Error
			if err != nil {
				return err
			}
		}
		err = tx.CreateInBatches(generateModelInventoryTransactions, len(generateModelInventoryTransactions)).Error
		if err != nil {
			return err
		}

		err = tx.Model(&models.SaleList{}).Where("id = ?", c.OrderId).Update("shipment_status", "1").Error
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("库存类型不合法")
}

// Update 修改InventoryTransactions对象
func (e *InventoryTransactions) Update(c *dto.InventoryTransactionsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.InventoryTransactions{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("InventoryTransactionsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除InventoryTransactions
func (e *InventoryTransactions) Remove(d *dto.InventoryTransactionsDeleteReq, p *actions.DataPermission) error {
	var data models.InventoryTransactions

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveInventoryTransactions error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
