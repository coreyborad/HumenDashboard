import {
  getUserStock,
  createUserStock,
  deleteUserStock,
  updateUserStock,
  getStockList
} from '@/api/stock'

const getDefaultState = () => {
  return {
    list: [],
    stock_list: []
  }
}

const state = getDefaultState()
const getters = {
}

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_LIST: (state, list) => {
    state.list = list
  },
  SET_STOCK_LIST: (state, list) => {
    state.stock_list = list
  },
}

const actions = {
  getList({ commit }) {
    return new Promise((resolve, reject) => {
      getUserStock().then(response => {
        const data = response
        commit('SET_LIST', data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  createUserStockInfo({ commit, dispatch }, data) {
    return new Promise((resolve, reject) => {
      createUserStock(data).then(response => {
        dispatch('getList')
        resolve(response.data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  deleteUserStockInfo({ commit, dispatch }, id) {
    return new Promise((resolve, reject) => {
      deleteUserStock(id).then(response => {
        dispatch('getList')
        resolve(response.data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  updateUserStock({ commit, dispatch }, form) {
    return new Promise((resolve, reject) => {
      updateUserStock(form).then(response => {
        dispatch('getList')
        resolve(response.data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  getStockList({ commit }) {
    return new Promise((resolve, reject) => {
      getStockList().then(response => {
        const data = response
        commit('SET_STOCK_LIST', data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}
