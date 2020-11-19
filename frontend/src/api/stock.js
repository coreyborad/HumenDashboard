import request from '@/utils/request'

export const getUserStock = async() => await request.get('user_stock')

export const createUserStock = async(data) => await request.post('user_stock', data)

export const deleteUserStock = async(id) => await request.delete('user_stock/' + id)

export const getStockList = async() => await request.get('stock')

// export function getInfo(token) {
//   return request({
//     url: '/vue-admin-template/user/info',
//     method: 'get',
//     params: { token }
//   })
// }

// export function logout() {
//   return request({
//     url: '/vue-admin-template/user/logout',
//     method: 'post'
//   })
// }
