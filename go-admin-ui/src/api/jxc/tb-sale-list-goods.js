import request from '@/utils/request'

// 查询SaleListGoods列表
export function listSaleListGoods(query) {
    return request({
        url: '/api/v1/sale-list-goods',
        method: 'get',
        params: query
    })
}

// 查询SaleListGoods详细
export function getSaleListGoods (id) {
    return request({
        url: '/api/v1/sale-list-goods/' + id,
        method: 'get'
    })
}


// 新增SaleListGoods
export function addSaleListGoods(data) {
    return request({
        url: '/api/v1/sale-list-goods',
        method: 'post',
        data: data
    })
}

// 修改SaleListGoods
export function updateSaleListGoods(data) {
    return request({
        url: '/api/v1/sale-list-goods/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SaleListGoods
export function delSaleListGoods(data) {
    return request({
        url: '/api/v1/sale-list-goods',
        method: 'delete',
        data: data
    })
}

