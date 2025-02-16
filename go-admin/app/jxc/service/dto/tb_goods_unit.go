package dto

import (

	"go-admin/app/jxc/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type GoodsUnitGetPageReq struct {
	dto.Pagination     `search:"-"`
    Name string `form:"name"  search:"type:exact;column:name;table:tb_goods_unit" comment:"商品单位"`
    GoodsUnitOrder
}

type GoodsUnitOrder struct {
    Id string `form:"idOrder"  search:"type:order;column:id;table:tb_goods_unit"`
    Name string `form:"nameOrder"  search:"type:order;column:name;table:tb_goods_unit"`
    CreatedAt string `form:"createdAtOrder"  search:"type:order;column:created_at;table:tb_goods_unit"`
    UpdatedAt string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:tb_goods_unit"`
    DeletedAt string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:tb_goods_unit"`
    CreateBy string `form:"createByOrder"  search:"type:order;column:create_by;table:tb_goods_unit"`
    UpdateBy string `form:"updateByOrder"  search:"type:order;column:update_by;table:tb_goods_unit"`
    
}

func (m *GoodsUnitGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type GoodsUnitInsertReq struct {
    Id int `json:"-" comment:"主键编码"` // 主键编码
    Name string `json:"name" comment:"商品单位"`
    common.ControlBy
}

func (s *GoodsUnitInsertReq) Generate(model *models.GoodsUnit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *GoodsUnitInsertReq) GetId() interface{} {
	return s.Id
}

type GoodsUnitUpdateReq struct {
    Id int `uri:"id" comment:"主键编码"` // 主键编码
    Name string `json:"name" comment:"商品单位"`
    common.ControlBy
}

func (s *GoodsUnitUpdateReq) Generate(model *models.GoodsUnit)  {
    if s.Id == 0 {
        model.Model = common.Model{ Id: s.Id }
    }
    model.Name = s.Name
    model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *GoodsUnitUpdateReq) GetId() interface{} {
	return s.Id
}

// GoodsUnitGetReq 功能获取请求参数
type GoodsUnitGetReq struct {
     Id int `uri:"id"`
}
func (s *GoodsUnitGetReq) GetId() interface{} {
	return s.Id
}

// GoodsUnitDeleteReq 功能删除请求参数
type GoodsUnitDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *GoodsUnitDeleteReq) GetId() interface{} {
	return s.Ids
}
