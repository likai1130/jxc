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

type GoodsUnit struct {
	api.Api
}

// GetPage 获取单位信息列表
// @Summary 获取单位信息列表
// @Description 获取单位信息列表
// @Tags 单位信息
// @Param name query string false "商品单位"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.GoodsUnit}} "{"code": 200, "data": [...]}"
// @Router /api/v1/goods-unit [get]
// @Security Bearer
func (e GoodsUnit) GetPage(c *gin.Context) {
    req := dto.GoodsUnitGetPageReq{}
    s := service.GoodsUnit{}
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
	list := make([]models.GoodsUnit, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取单位信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取单位信息
// @Summary 获取单位信息
// @Description 获取单位信息
// @Tags 单位信息
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.GoodsUnit} "{"code": 200, "data": [...]}"
// @Router /api/v1/goods-unit/{id} [get]
// @Security Bearer
func (e GoodsUnit) Get(c *gin.Context) {
	req := dto.GoodsUnitGetReq{}
	s := service.GoodsUnit{}
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
	var object models.GoodsUnit

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取单位信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建单位信息
// @Summary 创建单位信息
// @Description 创建单位信息
// @Tags 单位信息
// @Accept application/json
// @Product application/json
// @Param data body dto.GoodsUnitInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/goods-unit [post]
// @Security Bearer
func (e GoodsUnit) Insert(c *gin.Context) {
    req := dto.GoodsUnitInsertReq{}
    s := service.GoodsUnit{}
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
		e.Error(500, err, fmt.Sprintf("创建单位信息失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改单位信息
// @Summary 修改单位信息
// @Description 修改单位信息
// @Tags 单位信息
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.GoodsUnitUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/goods-unit/{id} [put]
// @Security Bearer
func (e GoodsUnit) Update(c *gin.Context) {
    req := dto.GoodsUnitUpdateReq{}
    s := service.GoodsUnit{}
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
		e.Error(500, err, fmt.Sprintf("修改单位信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除单位信息
// @Summary 删除单位信息
// @Description 删除单位信息
// @Tags 单位信息
// @Param data body dto.GoodsUnitDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/goods-unit [delete]
// @Security Bearer
func (e GoodsUnit) Delete(c *gin.Context) {
    s := service.GoodsUnit{}
    req := dto.GoodsUnitDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除单位信息失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
