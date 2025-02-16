import request from '@/utils/request'

// 查询GoodsUnit列表
export function listGoodsUnit(query) {
    return request({
        url: '/api/v1/goods-unit',
        method: 'get',
        params: query
    })
}

// 查询GoodsUnit详细
export function getGoodsUnit (id) {
    return request({
        url: '/api/v1/goods-unit/' + id,
        method: 'get'
    })
}


// 新增GoodsUnit
export function addGoodsUnit(data) {
    return request({
        url: '/api/v1/goods-unit',
        method: 'post',
        data: data
    })
}

// 修改GoodsUnit
export function updateGoodsUnit(data) {
    return request({
        url: '/api/v1/goods-unit/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除GoodsUnit
export function delGoodsUnit(data) {
    return request({
        url: '/api/v1/goods-unit',
        method: 'delete',
        data: data
    })
}

