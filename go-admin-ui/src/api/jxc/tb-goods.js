import request from '@/utils/request'

// 查询Goods列表
export function listGoods(query) {
    return request({
        url: '/api/v1/goods',
        method: 'get',
        params: query
    })
}

// 查询Goods详细
export function getGoods (id) {
    return request({
        url: '/api/v1/goods/' + id,
        method: 'get'
    })
}


// 新增Goods
export function addGoods(data) {
    return request({
        url: '/api/v1/goods',
        method: 'post',
        data: data
    })
}

// 修改Goods
export function updateGoods(data) {
    return request({
        url: '/api/v1/goods/'+data.id,
        method: 'put',
        data: data
    })
}

// 删除Goods
export function delGoods(data) {
    return request({
        url: '/api/v1/goods',
        method: 'delete',
        data: data
    })
}

