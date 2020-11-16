import request from '@/utils/request'

// 有做Promise
export const getUserStock = async() => await request.get('user_stock')

export const createUserStock = async(data) => await request.post('user_stock', data)

export const deleteUserStock = async(id) => await request.delete('user_stock/' + id)
