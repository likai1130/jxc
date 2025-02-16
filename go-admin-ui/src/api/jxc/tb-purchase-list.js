import request from '@/utils/request'

// 查询PurchaseList列表
export function listPurchaseList(query) {
    return request({
        url: '/api/v1/purchase-list',
        method: 'get',
        params: query
    })
}

// 查询PurchaseList详细
export function getPurchaseList (id) {
    return request({
        url: '/api/v1/purchase-list/' + id,
        method: 'get'
    })
}


// 新增PurchaseList
export function addPurchaseList(data) {
    return request({
        url: '/api/v1/purchase-list',
        method: 'post',
        data: data
    })
}

// 修改PurchaseList
export function updatePurchaseList(data) {
    return request({
        url: '/api/v1/purchase-list/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除PurchaseList
export function delPurchaseList(data) {
    return request({
        url: '/api/v1/purchase-list',
        method: 'delete',
        data: data
    })
}

