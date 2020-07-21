import {
  // eslint-disable-next-line camelcase
  BundleService_ListSettingsBundles,
  // eslint-disable-next-line camelcase
  ValueService_SaveSettingsValue
} from '../client/settings'
import axios from 'axios'
import keyBy from 'lodash/keyBy'

const state = {
  config: null,
  initialized: false,
  settingsBundles: {}
}

const getters = {
  config: state => state.config,
  initialized: state => state.initialized,
  extensions: state => {
    return [...new Set(Object.values(state.settingsBundles).map(bundle => bundle.extension))].sort()
  },
  getBundlesByExtension: state => extension => {
    return Object.values(state.settingsBundles)
      .filter(bundle => bundle.extension === extension)
      .sort((b1, b2) => {
        return b1.name.localeCompare(b2.name)
      })
  }
}

const mutations = {
  SET_INITIALIZED (state, value) {
    state.initialized = value
  },
  SET_SETTINGS_BUNDLES (state, settingsBundles) {
    state.settingsBundles = keyBy(settingsBundles, 'id')
  },
  LOAD_CONFIG (state, config) {
    state.config = config
  }
}

const actions = {
  // Used by ocis-web.
  loadConfig ({ commit }, config) {
    commit('LOAD_CONFIG', config)
  },

  async initialize ({ commit, dispatch }) {
    await dispatch('fetchBundles')
    commit('SET_INITIALIZED', true)
  },

  async fetchBundles ({ commit, dispatch, rootGetters }) {
    injectAuthToken(rootGetters)
    try {
      const response = await BundleService_ListSettingsBundles({
        $domain: rootGetters.configuration.server,
        body: {
          accountUuid: 'me'
        }
      })
      if (response.status === 201) {
        // the settings markup has implicit typing. inject an explicit type variable here
        const bundles = response.data.bundles
        if (bundles) {
          bundles.forEach(bundle => {
            bundle.settings.forEach(setting => {
              if (setting.intValue) {
                setting.type = 'number'
              } else if (setting.stringValue) {
                setting.type = 'string'
              } else if (setting.boolValue) {
                setting.type = 'boolean'
              } else if (setting.singleChoiceValue) {
                setting.type = 'singleChoice'
              } else if (setting.multiChoiceValue) {
                setting.type = 'multiChoice'
              } else {
                setting.type = 'unknown'
              }
            })
          })
          commit('SET_SETTINGS_BUNDLES', bundles)
        } else {
          commit('SET_SETTINGS_BUNDLES', [])
        }
      }
    } catch (err) {
      dispatch('showMessage', {
        title: 'Failed to fetch settings bundles.',
        status: 'danger'
      }, { root: true })
    }
  },

  async saveValue ({ commit, dispatch, getters, rootGetters }, { bundle, setting, payload }) {
    injectAuthToken(rootGetters)
    try {
      const response = await ValueService_SaveSettingsValue({
        $domain: rootGetters.configuration.server,
        body: {
          value: payload
        }
      })
      if (response.status === 201 && response.data.value) {
        commit('SET_SETTINGS_VALUE', response.data.value, { root: true })
      }
    } catch (e) {
      dispatch('showMessage', {
        title: `Failed to save »${setting.displayName}«.`,
        status: 'danger'
      }, { root: true })
    }
  }
}

export default {
  namespaced: true,
  state,
  getters,
  actions,
  mutations
}

function injectAuthToken (rootGetters) {
  axios.interceptors.request.use(config => {
    if (typeof config.headers.Authorization === 'undefined') {
      const token = rootGetters.user.token
      if (token) {
        config.headers.Authorization = `Bearer ${token}`
      }
    }
    return config
  })
}
