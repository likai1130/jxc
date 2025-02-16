import request from '@/utils/request'

// 查询Supplier列表
export function listSupplier(query) {
    return request({
        url: '/api/v1/supplier',
        method: 'get',
        params: query
    })
}

// 查询Supplier详细
export function getSupplier (id) {
    return request({
        url: '/api/v1/supplier/' + id,
        method: 'get'
    })
}


// 新增Supplier
export function addSupplier(data) {
    return request({
        url: '/api/v1/supplier',
        method: 'post',
        data: data
    })
}

// 修改Supplier
export function updateSupplier(data) {
    return request({
        url: '/api/v1/supplier/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除Supplier
export function delSupplier(data) {
    return request({
        url: '/api/v1/supplier',
        method: 'delete',
        data: data
    })
}

