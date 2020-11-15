import request from '@/utils/request'

// 有做Promise
export const getStock = async() => await request.get('stock')
