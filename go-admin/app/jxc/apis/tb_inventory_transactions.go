package apis

import (
    "fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/jxc/models"
	"go-admin/app/jxc/service"
	"go-admin/app/jxc/service/dto"
	"go-admin/common/actions"
)

type InventoryTransactions struct {
	api.Api
}

// GetPage 获取库存管理列表
// @Summary 获取库存管理列表
// @Description 获取库存管理列表
// @Tags 库存管理
// @Param transactionType query string false "操作类型"
// @Param gCode query string false "商品编码"
// @Param gName query string false "商品名称"
// @Param orderNumber query string false "订单号"
// @Param goodsId query int64 false "商品ID"
// @Param orderId query int64 false "订单ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.InventoryTransactions}} "{"code": 200, "data": [...]}"
// @Router /api/v1/inventory-transactions [get]
// @Security Bearer
func (e InventoryTransactions) GetPage(c *gin.Context) {
    req := dto.InventoryTransactionsGetPageReq{}
    s := service.InventoryTransactions{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
   	if err != nil {
   		e.Logger.Error(err)
   		e.Error(500, err, err.Error())
   		return
   	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.InventoryTransactions, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取库存管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取库存管理
// @Summary 获取库存管理
// @Description 获取库存管理
// @Tags 库存管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.InventoryTransactions} "{"code": 200, "data": [...]}"
// @Router /api/v1/inventory-transactions/{id} [get]
// @Security Bearer
func (e InventoryTransactions) Get(c *gin.Context) {
	req := dto.InventoryTransactionsGetReq{}
	s := service.InventoryTransactions{}
    err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.InventoryTransactions

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取库存管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建库存管理
// @Summary 创建库存管理
// @Description 创建库存管理
// @Tags 库存管理
// @Accept application/json
// @Product application/json
// @Param data body dto.InventoryTransactionsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/inventory-transactions [post]
// @Security Bearer
func (e InventoryTransactions) Insert(c *gin.Context) {
    req := dto.InventoryTransactionsInsertReq{}
    s := service.InventoryTransactions{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建库存管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改库存管理
// @Summary 修改库存管理
// @Description 修改库存管理
// @Tags 库存管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.InventoryTransactionsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/inventory-transactions/{id} [put]
// @Security Bearer
func (e InventoryTransactions) Update(c *gin.Context) {
    req := dto.InventoryTransactionsUpdateReq{}
    s := service.InventoryTransactions{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改库存管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除库存管理
// @Summary 删除库存管理
// @Description 删除库存管理
// @Tags 库存管理
// @Param data body dto.InventoryTransactionsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/inventory-transactions [delete]
// @Security Bearer
func (e InventoryTransactions) Delete(c *gin.Context) {
    s := service.InventoryTransactions{}
    req := dto.InventoryTransactionsDeleteReq{}
    err := e.MakeContext(c).
        MakeOrm().
        Bind(&req).
        MakeService(&s.Service).
        Errors
    if err != nil {
        e.Logger.Error(err)
        e.Error(500, err, err.Error())
        return
    }

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除库存管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
