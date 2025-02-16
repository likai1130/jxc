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

type PurchaseList struct {
	api.Api
}

// GetPage 获取查询采购单列表
// @Summary 获取查询采购单列表
// @Description 获取查询采购单列表
// @Tags 查询采购单
// @Param purchaseNumber query string false "采购单号"
// @Param state query string false "是否付款"
// @Param storageStatus query string false "是否入库"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.PurchaseList}} "{"code": 200, "data": [...]}"
// @Router /api/v1/purchase-list [get]
// @Security Bearer
func (e PurchaseList) GetPage(c *gin.Context) {
    req := dto.PurchaseListGetPageReq{}
    s := service.PurchaseList{}
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
	list := make([]models.PurchaseList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取查询采购单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取查询采购单
// @Summary 获取查询采购单
// @Description 获取查询采购单
// @Tags 查询采购单
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.PurchaseList} "{"code": 200, "data": [...]}"
// @Router /api/v1/purchase-list/{id} [get]
// @Security Bearer
func (e PurchaseList) Get(c *gin.Context) {
	req := dto.PurchaseListGetReq{}
	s := service.PurchaseList{}
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
	var object models.PurchaseList

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取查询采购单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建查询采购单
// @Summary 创建查询采购单
// @Description 创建查询采购单
// @Tags 查询采购单
// @Accept application/json
// @Product application/json
// @Param data body dto.PurchaseListInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/purchase-list [post]
// @Security Bearer
func (e PurchaseList) Insert(c *gin.Context) {
    req := dto.PurchaseListInsertReq{}
    s := service.PurchaseList{}
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
		e.Error(500, err, fmt.Sprintf("创建查询采购单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改查询采购单
// @Summary 修改查询采购单
// @Description 修改查询采购单
// @Tags 查询采购单
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.PurchaseListUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/purchase-list/{id} [put]
// @Security Bearer
func (e PurchaseList) Update(c *gin.Context) {
    req := dto.PurchaseListUpdateReq{}
    s := service.PurchaseList{}
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
		e.Error(500, err, fmt.Sprintf("修改查询采购单失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除查询采购单
// @Summary 删除查询采购单
// @Description 删除查询采购单
// @Tags 查询采购单
// @Param data body dto.PurchaseListDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/purchase-list [delete]
// @Security Bearer
func (e PurchaseList) Delete(c *gin.Context) {
    s := service.PurchaseList{}
    req := dto.PurchaseListDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除查询采购单失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
