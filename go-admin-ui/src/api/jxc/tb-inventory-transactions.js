import request from '@/utils/request'

// 查询InventoryTransactions列表
export function listInventoryTransactions(query) {
    return request({
        url: '/api/v1/inventory-transactions',
        method: 'get',
        params: query
    })
}

// 查询InventoryTransactions详细
export function getInventoryTransactions (id) {
    return request({
        url: '/api/v1/inventory-transactions/' + id,
        method: 'get'
    })
}


// 新增InventoryTransactions
export function addInventoryTransactions(data) {
    return request({
        url: '/api/v1/inventory-transactions',
        method: 'post',
        data: data
    })
}

// 修改InventoryTransactions
export function updateInventoryTransactions(data) {
    return request({
        url: '/api/v1/inventory-transactions/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除InventoryTransactions
export function delInventoryTransactions(data) {
    return request({
        url: '/api/v1/inventory-transactions',
        method: 'delete',
        data: data
    })
}

