import axios from 'axios'
import store from '@/store'
import { getToken } from '@/utils/auth'

console.log(process.env, 'asdasdasd')
// create an axios instance
const service = axios.create({
  baseURL: process.env.VUE_APP_WEB_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 30000 // request timeout
})

// request interceptor
service.interceptors.request.use(
  config => {
    // do something before request is sent
    if (store.getters.token) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['Authorization'] = 'Bearer ' + getToken()
    }
    return config
  },
  error => {
    // do something with request error
    return Promise.reject(error)
  }
)

let refreshing = false
let requests = []

service.interceptors.response.use(
  response => response.data,
  async error => {
    const { response: errorResponse } = error

    if (refreshing) {
      return new Promise((resolve) => {
        requests.push(() => {
          resolve(service(errorResponse.config))
        })
      })
    } else {
      if (errorResponse && errorResponse.data && errorResponse.data.error.indexOf('token is expired') > -1) {
        try {
          await store.dispatch('user/refreshToken', getToken())

          requests.forEach(cb => cb())
          requests = []

          return service(errorResponse.config)
        } catch (error) {
          location.reload()
        } finally {
          refreshing = false
        }
      }
    }
    return Promise.reject(error)
  }
)

export default service

