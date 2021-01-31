import request from '@/utils/request'

export const getUserStock = async() => await request.get('user_stock')

export const createUserStock = async(data) => await request.post('user_stock', data)

export const deleteUserStock = async(id) => await request.delete('user_stock/' + id)

export const getStockList = async() => await request.get('stock')

export const updateUserStock = async (data) => await request.patch('user_stock/' + data.id, data)
