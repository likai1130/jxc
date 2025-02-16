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

type SaleListGoods struct {
	api.Api
}

// GetPage 获取创建销售单列表
// @Summary 获取创建销售单列表
// @Description 获取创建销售单列表
// @Tags 创建销售单
// @Param gCode query string false "商品编码"
// @Param gName query string false "商品名称"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.SaleListGoods}} "{"code": 200, "data": [...]}"
// @Router /api/v1/sale-list-goods [get]
// @Security Bearer
func (e SaleListGoods) GetPage(c *gin.Context) {
    req := dto.SaleListGoodsGetPageReq{}
    s := service.SaleListGoods{}
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
	list := make([]models.SaleListGoods, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取创建销售单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取创建销售单
// @Summary 获取创建销售单
// @Description 获取创建销售单
// @Tags 创建销售单
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.SaleListGoods} "{"code": 200, "data": [...]}"
// @Router /api/v1/sale-list-goods/{id} [get]
// @Security Bearer
func (e SaleListGoods) Get(c *gin.Context) {
	req := dto.SaleListGoodsGetReq{}
	s := service.SaleListGoods{}
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
	var object models.SaleListGoods

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取创建销售单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK( object, "查询成功")
}

// Insert 创建创建销售单
// @Summary 创建创建销售单
// @Description 创建创建销售单
// @Tags 创建销售单
// @Accept application/json
// @Product application/json
// @Param data body dto.SaleListGoodsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/sale-list-goods [post]
// @Security Bearer
func (e SaleListGoods) Insert(c *gin.Context) {
    req := dto.SaleListGoodsInsertReq{}
    s := service.SaleListGoods{}
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
		e.Error(500, err, fmt.Sprintf("创建创建销售单失败，\r\n失败信息 %s", err.Error()))
        return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改创建销售单
// @Summary 修改创建销售单
// @Description 修改创建销售单
// @Tags 创建销售单
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.SaleListGoodsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/sale-list-goods/{id} [put]
// @Security Bearer
func (e SaleListGoods) Update(c *gin.Context) {
    req := dto.SaleListGoodsUpdateReq{}
    s := service.SaleListGoods{}
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
		e.Error(500, err, fmt.Sprintf("修改创建销售单失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "修改成功")
}

// Delete 删除创建销售单
// @Summary 删除创建销售单
// @Description 删除创建销售单
// @Tags 创建销售单
// @Param data body dto.SaleListGoodsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/sale-list-goods [delete]
// @Security Bearer
func (e SaleListGoods) Delete(c *gin.Context) {
    s := service.SaleListGoods{}
    req := dto.SaleListGoodsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除创建销售单失败，\r\n失败信息 %s", err.Error()))
        return
	}
	e.OK( req.GetId(), "删除成功")
}
