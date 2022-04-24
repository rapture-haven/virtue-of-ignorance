import Vue from "vue";
import Vuex from "vuex";
import { get, set } from '~/utils/localstorage'
import { useAccessor, mutationTree, actionTree } from 'typed-vuex'

export const state = () => ({
  displayname: get<string>('displayname', ''),
  password: get<string>('password', ''),
  active: false,
  connecting: false,
  connected: false,
  locked: {} as Record<string, boolean>,
})

export const storePattern = {
  state: state,
  mutations: {},
  actions: {},
  modules: {},
}

Vue.use(Vuex);

const store = new Vuex.Store(storePattern)
export const accessor = useAccessor(store, storePattern)