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

type Consumer struct {
	api.Api
}

// GetPage 获取客户信息管理列表
// @Summary 获取客户信息管理列表
// @Description 获取客户信息管理列表
// @Tags 客户信息管理
// @Param customerName query string false "客户名称"
// @Param customerType query string false "客户类型"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Consumer}} "{"code": 200, "data": [...]}"
// @Router /api/v1/consumer [get]
// @Security Bearer
func (e Consumer) GetPage(c *gin.Context) {
    req := dto.ConsumerGetPageReq{}
    s := service.Consumer{}
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
	list := make([]models.Consumer, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取客户信息管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取客户信息管理
// @Summary 获取客户信息管理
// @Description 获取客户信息管理
// @Tags 客户信息管理
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Consumer} "{"code": 200, "data": [...]}"
// @Router /api/v1/consumer/{id} [get]
// @Security Bearer
func (e Consumer) Get(c *gin.Context) {
	req := dto.ConsumerGetReq{}
	s := service.Consumer{}
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
	var object models.Consumer

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取客户信息管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建客户信息管理
// @Summary 创建客户信息管理
// @Description 创建客户信息管理
// @Tags 客户信息管理
// @Accept application/json
// @Product application/json
// @Param data body dto.ConsumerInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/consumer [post]
// @Security Bearer
func (e Consumer) Insert(c *gin.Context) {
    req := dto.ConsumerInsertReq{}
    s := service.Consumer{}
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
		e.Error(500, err, fmt.Sprintf("创建客户信息管理失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改客户信息管理
// @Summary 修改客户信息管理
// @Description 修改客户信息管理
// @Tags 客户信息管理
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ConsumerUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/consumer/{id} [put]
// @Security Bearer
func (e Consumer) Update(c *gin.Context) {
    req := dto.ConsumerUpdateReq{}
    s := service.Consumer{}
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
		e.Error(500, err, fmt.Sprintf("修改客户信息管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除客户信息管理
// @Summary 删除客户信息管理
// @Description 删除客户信息管理
// @Tags 客户信息管理
// @Param data body dto.ConsumerDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/consumer [delete]
// @Security Bearer
func (e Consumer) Delete(c *gin.Context) {
    s := service.Consumer{}
    req := dto.ConsumerDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除客户信息管理失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
