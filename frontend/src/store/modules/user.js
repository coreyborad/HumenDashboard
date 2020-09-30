import { login, logout, getUserInfo } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'

const state = {
  token: getToken(),
  id: 0,
  name: '',
  email: ''
}

const mutations = {
  SET_USER: (state, user) => {
    state.id = user.id
    state.name = user.name
    state.email = user.email
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  }
}

const actions = {
  // user login
  async login({ commit }, loginInfo) {
    const { email, password } = loginInfo
    try {
      const { access_token } = await login(
        {
          email: email.trim(),
          password: password.trim()
        }
      )
      setToken(access_token)
    } catch (error) {
      throw error
    }
  },
  // user logout
  async logout({ commit, state }) {
    try {
      await logout()
      removeToken()
    } catch (error) {
      throw error
    }
  },
  // set user
  async setUser({ commit }) {
    try {
      const data = await getUserInfo()
      commit('SET_USER', data)
    } catch (error) {
      throw error
    }
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

