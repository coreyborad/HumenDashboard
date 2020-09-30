import request from '@/utils/request'

// 有做Promise
export const login = async data =>
  await request.post('login', {
    grant_type: 'password',
    email: data.email,
    password: data.password
  })
