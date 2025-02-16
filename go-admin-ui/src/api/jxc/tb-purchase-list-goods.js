import request from '@/utils/request'

// 查询PurchaseListGoods列表
export function listPurchaseListGoods(query) {
    return request({
        url: '/api/v1/purchase-list-goods',
        method: 'get',
        params: query
    })
}

// 查询PurchaseListGoods详细
export function getPurchaseListGoods (id) {
    return request({
        url: '/api/v1/purchase-list-goods/' + id,
        method: 'get'
    })
}


// 新增PurchaseListGoods
export function addPurchaseListGoods(data) {
    return request({
        url: '/api/v1/purchase-list-goods',
        method: 'post',
        data: data
    })
}

// 修改PurchaseListGoods
export function updatePurchaseListGoods(data) {
    return request({
        url: '/api/v1/purchase-list-goods/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除PurchaseListGoods
export function delPurchaseListGoods(data) {
    return request({
        url: '/api/v1/purchase-list-goods',
        method: 'delete',
        data: data
    })
}

