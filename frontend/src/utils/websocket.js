import qs from 'qs'
import {
  getToken
  // tokenExpired
} from '@/utils/auth'
import ReconnectingWebsocket from 'reconnectingwebsocket'
// import store from '@/store'

export const connect = (path, query) => new Promise((resolve, reject) => {
  // if (tokenExpired()) {
  //   await store.dispatch('user/refreshToken', getToken())
  // }

  query = qs.stringify({
    ...query,
    access_token: getToken()
  })

  const ws = new ReconnectingWebsocket(`${process.env.VUE_APP_WEBSOCKET_URL}${path}?${query}`)

  ws.onopen = async() => {
    // if (tokenExpired()) {
    //   await store.dispatch('user/refreshToken', getToken())
    // }

    return resolve(ws)
  }
})

export default connect
