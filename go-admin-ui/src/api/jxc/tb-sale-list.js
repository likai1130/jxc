import request from '@/utils/request'

// 查询SaleList列表
export function listSaleList(query) {
    return request({
        url: '/api/v1/sale-list',
        method: 'get',
        params: query
    })
}

// 查询SaleList详细
export function getSaleList (id) {
    return request({
        url: '/api/v1/sale-list/' + id,
        method: 'get'
    })
}


// 新增SaleList
export function addSaleList(data) {
    return request({
        url: '/api/v1/sale-list',
        method: 'post',
        data: data
    })
}

// 修改SaleList
export function updateSaleList(data) {
    return request({
        url: '/api/v1/sale-list/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除SaleList
export function delSaleList(data) {
    return request({
        url: '/api/v1/sale-list',
        method: 'delete',
        data: data
    })
}

