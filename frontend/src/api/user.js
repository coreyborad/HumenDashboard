import request from '@/utils/request'

export const login = async data =>
  await request.post('login', {
    grant_type: 'password',
    email: data.email,
    password: data.password
  })

export const getInfo = async() =>
  await request.get('user')

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
