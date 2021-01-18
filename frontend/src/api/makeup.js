import request from '@/utils/request'

export const getMakeup = async() => await request.get('makeup/list')

export const getMakeupByQuery = async(query) => await request.get('makeup', { params: query })

export const createMakeupInfo = async(data) => await request.post('makeup', data)

export const deleteMakeupInfo = async(id) => await request.delete('makeup/' + id)

export const updateMakeupInfo = async(data) => await request.patch('makeup/' + data.id, data)

export const createMakeupCost = async(data) => await request.post('makeup/cost', data)

export const deleteMakeupCost = async (id) => await request.delete('makeup/cost/' + id)

export const updateMakeupCost = async(data) => await request.patch('makeup/cost/' + data.id, data)

export const createMakeupSale = async(data) => await request.post('makeup/sale', data)

export const deleteMakeupSale = async(id) => await request.delete('makeup/sale/' + id)

export const updateMakeupSale = async (data) => await request.patch('makeup/sale/' + data.id, data)
