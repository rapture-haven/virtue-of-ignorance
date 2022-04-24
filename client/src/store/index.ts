import Vue from "vue";
import Vuex from "vuex";
import { useAccessor, mutationTree, actionTree } from "typed-vuex";
import { get, set } from "@/utils/localstorage";
import { EVENT } from "@/virtue-of-ignorance/events";
import { AdminLockResource } from "@/virtue-of-ignorance/messages";

import * as video from './video'
import * as remote from './remote'
import * as user from './user'
import * as settings from './settings'
import * as client from './client'

export const state = () => ({
  displayname: get<string>("displayname", ""),
  password: get<string>("password", ""),
  active: false,
  connecting: false,
  connected: false,
  locked: {} as Record<string, boolean>,
});

export const mutations = mutationTree(state, {
  setActive(state) {
    state.active = true;
  },

  setLogin(
    state,
    { displayname, password }: { displayname: string; password: string }
  ) {
    state.displayname = displayname;
    state.password = password;
  },

  setLocked(state, resource: string) {
    Vue.set(state.locked, resource, true);
  },

  setUnlocked(state, resource: string) {
    Vue.set(state.locked, resource, false);
  },

  setConnnecting(state) {
    state.connected = false;
    state.connecting = true;
  },

  setConnected(state, connected: boolean) {
    state.connected = connected;
    state.connecting = false;
    if (connected) {
      set("displayname", state.displayname);
      set("password", state.password);
    }
  },
});

export const actions = actionTree(
  { state, mutations },
  {
    initialise(store) {
      accessor.settings.initialise()
    },

    lock(_, resource: AdminLockResource) {
      if (!accessor.connected || !accessor.user.admin) {
        return;
      }

      $client.sendMessage(EVENT.ADMIN.LOCK, { resource });
    },

    unlock(_, resource: AdminLockResource) {
      if (!accessor.connected || !accessor.user.admin) {
        return;
      }

      $client.sendMessage(EVENT.ADMIN.UNLOCK, { resource });
    },

    login(
      { state },
      { displayname, password }: { displayname: string; password: string }
    ) {
      console.log(displayname, password);
      accessor.setLogin({ displayname, password });
      $client.login(password, displayname);
    },

    logout({ state }) {
      accessor.setLogin({ displayname: "", password: "" });
      set("displayname", "");
      set("password", "");
      $client.logout();
    },
  }
);

export const storePattern = {
  state: state,
  mutations: mutations,
  actions: actions,
  modules: { video, user, remote, settings, client },
};

Vue.use(Vuex);

const store = new Vuex.Store(storePattern);
export const accessor = useAccessor(store, storePattern);

Vue.prototype.$accessor = accessor;
declare module "vue/types/vue" {
  interface Vue {
    $accessor: typeof accessor;
  }
}

export default store;
