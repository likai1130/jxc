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

type SaleList struct {
	api.Api
}

// GetPage 获取销售管理列表
// @Summary 获取销售管理列表
// @Description 获取销售管理列表
// @Tags 销售管理
// @Param saleNumber query string false "销售单号"
// @Param state query string false "状态"
// @Param customerId query int64 false "消费者ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SaleList}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sale-list [get]
// @Security Bearer
func (e SaleList) GetPage(c *gin.Context) {
    req := dto.SaleListGetPageReq{}
    s := service.SaleList{}
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
	list := make([]models.SaleList, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取销售管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取销售管理
// @Summary 获取销售管理
// @Description 获取销售管理
// @Tags 销售管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SaleList} "{"code": 200, "data": [...]}"
// @Router /api/v1/sale-list/{id} [get]
// @Security Bearer
func (e SaleList) Get(c *gin.Context) {
	req := dto.SaleListGetReq{}
	s := service.SaleList{}
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
	var object models.SaleList

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取销售管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建销售管理
// @Summary 创建销售管理
// @Description 创建销售管理
// @Tags 销售管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SaleListInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sale-list [post]
// @Security Bearer
func (e SaleList) Insert(c *gin.Context) {
    req := dto.SaleListInsertReq{}
    s := service.SaleList{}
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
		e.Error(500, err, fmt.Sprintf("创建销售管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改销售管理
// @Summary 修改销售管理
// @Description 修改销售管理
// @Tags 销售管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SaleListUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sale-list/{id} [put]
// @Security Bearer
func (e SaleList) Update(c *gin.Context) {
    req := dto.SaleListUpdateReq{}
    s := service.SaleList{}
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
		e.Error(500, err, fmt.Sprintf("修改销售管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除销售管理
// @Summary 删除销售管理
// @Description 删除销售管理
// @Tags 销售管理
// @Param data body dto.SaleListDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sale-list [delete]
// @Security Bearer
func (e SaleList) Delete(c *gin.Context) {
    s := service.SaleList{}
    req := dto.SaleListDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除销售管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
