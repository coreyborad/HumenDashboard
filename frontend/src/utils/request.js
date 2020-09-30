import axios from 'axios'
import store from '@/store'
import { getToken } from '@/utils/auth'

// create an axios instance
const service = axios.create({
  baseURL: process.env.WEB_API, // url = base url + request url
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

const mockservice = axios.create({
  baseURL: process.env.VUE_APP_MOCK_API, // url = base url + request url
  // withCredentials: true, // send cookies when cross-domain requests
  timeout: 5000 // request timeout
})

// request interceptor
mockservice.interceptors.request.use(
  config => {
    // do something before request is sent

    if (store.getters.token) {
      // let each request carry token
      // ['X-Token'] is a custom headers key
      // please modify it according to the actual situation
      config.headers['X-Token'] = getToken()
    }
    return config
  },
  error => {
    // do something with request error
    console.log(error) // for debug
    return Promise.reject(error)
  }
)

// response interceptor
mockservice.interceptors.response.use(
  /**
   * If you want to get http information such as headers or status
   * Please return  response => response
  */

  /**
   * Determine the request status by custom code
   * Here is just an example
   * You can also judge the status by HTTP Status Code
   */
  response => {
    const res = response.data

    // if the custom code is not 20000, it is judged as an error.
    if (res.code !== 20000) {
      // 50008: Illegal token; 50012: Other clients logged in; 50014: Token expired;
      if (res.code === 50008 || res.code === 50012 || res.code === 50014) {
        console.log('mock')
      }
      return Promise.reject(new Error(res.message || 'Error'))
    } else {
      return res
    }
  },
  error => {
    return Promise.reject(error)
  }
)

export { mockservice, service }

export default service

