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

type Supplier struct {
	api.Api
}

// GetPage 获取供应商管理列表
// @Summary 获取供应商管理列表
// @Description 获取供应商管理列表
// @Tags 供应商管理
// @Param supplierName query string false "供应商名称"
// @Param contactName query string false "联系人名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Supplier}} "{"code": 200, "data": [...]}"
// @Router /api/v1/supplier [get]
// @Security Bearer
func (e Supplier) GetPage(c *gin.Context) {
    req := dto.SupplierGetPageReq{}
    s := service.Supplier{}
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
	list := make([]models.Supplier, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取供应商管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取供应商管理
// @Summary 获取供应商管理
// @Description 获取供应商管理
// @Tags 供应商管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Supplier} "{"code": 200, "data": [...]}"
// @Router /api/v1/supplier/{id} [get]
// @Security Bearer
func (e Supplier) Get(c *gin.Context) {
	req := dto.SupplierGetReq{}
	s := service.Supplier{}
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
	var object models.Supplier

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取供应商管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建供应商管理
// @Summary 创建供应商管理
// @Description 创建供应商管理
// @Tags 供应商管理
// @Accept application/json
// @Product application/json
// @Param data body dto.SupplierInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/supplier [post]
// @Security Bearer
func (e Supplier) Insert(c *gin.Context) {
    req := dto.SupplierInsertReq{}
    s := service.Supplier{}
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
		e.Error(500, err, fmt.Sprintf("创建供应商管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改供应商管理
// @Summary 修改供应商管理
// @Description 修改供应商管理
// @Tags 供应商管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SupplierUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/supplier/{id} [put]
// @Security Bearer
func (e Supplier) Update(c *gin.Context) {
    req := dto.SupplierUpdateReq{}
    s := service.Supplier{}
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
		e.Error(500, err, fmt.Sprintf("修改供应商管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除供应商管理
// @Summary 删除供应商管理
// @Description 删除供应商管理
// @Tags 供应商管理
// @Param data body dto.SupplierDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/supplier [delete]
// @Security Bearer
func (e Supplier) Delete(c *gin.Context) {
    s := service.Supplier{}
    req := dto.SupplierDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除供应商管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
