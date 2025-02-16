import request from '@/utils/request'

// 查询Consumer列表
export function listConsumer(query) {
    return request({
        url: '/api/v1/consumer',
        method: 'get',
        params: query
    })
}

// 查询Consumer详细
export function getConsumer (id) {
    return request({
        url: '/api/v1/consumer/' + id,
        method: 'get'
    })
}


// 新增Consumer
export function addConsumer(data) {
    return request({
        url: '/api/v1/consumer',
        method: 'post',
        data: data
    })
}

// 修改Consumer
export function updateConsumer(data) {
    return request({
        url: '/api/v1/consumer/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除Consumer
export function delConsumer(data) {
    return request({
        url: '/api/v1/consumer',
        method: 'delete',
        data: data
    })
}

