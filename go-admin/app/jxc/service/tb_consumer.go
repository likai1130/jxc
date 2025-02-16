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

type Consumer struct {
	service.Service
}

// GetPage 获取Consumer列表
func (e *Consumer) GetPage(c *dto.ConsumerGetPageReq, p *actions.DataPermission, list *[]models.Consumer, count *int64) error {
	var err error
	var data models.Consumer

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ConsumerService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Consumer对象
func (e *Consumer) Get(d *dto.ConsumerGetReq, p *actions.DataPermission, model *models.Consumer) error {
	var data models.Consumer

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetConsumer error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Consumer对象
func (e *Consumer) Insert(c *dto.ConsumerInsertReq) error {
    var err error
    var data models.Consumer
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ConsumerService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改Consumer对象
func (e *Consumer) Update(c *dto.ConsumerUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.Consumer{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("ConsumerService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除Consumer
func (e *Consumer) Remove(d *dto.ConsumerDeleteReq, p *actions.DataPermission) error {
	var data models.Consumer

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveConsumer error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
