import { getMakeup, createMakeupInfo, deleteMakeupInfo, updateMakeupInfo } from '@/api/makeup'

const getDefaultState = () => {
  return {
    list: [],
    brand: '',
    name: '',
    color_name: '',
    id: 0
  }
}

const state = getDefaultState()
const getters = {
  colorList: state => () => {
    if (state.brand === '' || state.name === '') {
      return []
    }
    return state.list.find(l => {
      if (l.brand === state.brand && l.name === state.name) {
        return true
      }
    }).color_list
  },
  costList: (state, getters) => () => {
    if (state.color_name === '') {
      return []
    }
    return getters.colorList.find(l => {
      if (l.color_name === state.color_name) {
        return true
      }
    }).costs
  },
  saleList: (state, getters) => () => {
    if (state.color_name === '') {
      return []
    }
    return getters.colorList.find(l => {
      if (l.color_name === state.color_name) {
        return true
      }
    }).sales
  }
}

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_LIST: (state, list) => {
    state.list = list
  },
  SET_BRAND: (state, brand) => {
    state.brand = brand
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_COLORNAME: (state, color_name) => {
    state.color_name = color_name
  },
  SET_ID: (state, id) => {
    state.id = id
  }
}

const actions = {
  getList({ commit }) {
    return new Promise((resolve, reject) => {
      getMakeup().then(response => {
        const data = response
        commit('SET_LIST', data)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  createMakeupInfo({ commit, dispatch }, data) {
    return new Promise((resolve, reject) => {
      createMakeupInfo(data).then(response => {
        dispatch('getList')
        resolve(response.data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  deleteMakeupInfo({ commit, dispatch }, id) {
    return new Promise((resolve, reject) => {
      deleteMakeupInfo(id).then(response => {
        dispatch('getList')
        resolve(response.data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  updateMakeupInfo({ commit, dispatch }, form) {
    return new Promise((resolve, reject) => {
      updateMakeupInfo(form).then(response => {
        dispatch('getList')
        resolve(response.data)
      }).catch(error => {
        reject(error)
      })
    })
  },
  setCurrentValue({ commit }, data) {
    return new Promise((resolve, reject) => {
      switch (data.target) {
        case 'brand':
          commit('SET_BRAND', data.value)
          break
        case 'name':
          commit('SET_NAME', data.value)
          break
        case 'color_name':
          commit('SET_COLORNAME', data.value)
          break
        case 'id':
          commit('SET_ID', data.value)
          break
      }
      resolve(data)
    })
  }
}

export default {
  namespaced: true,
  state,
  getters,
  mutations,
  actions
}

