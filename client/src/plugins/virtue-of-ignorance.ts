import { PluginObject } from "vue";
import { VirtueClient } from "@/virtue-of-ignorance";

declare global {
  const $client: VirtueClient;

  interface Window {
    $client: VirtueClient;
  }
}

declare module "vue/types/vue" {
  interface Vue {
    $client: VirtueClient;
  }
}

const plugin: PluginObject<undefined> = {
  install(Vue) {
    window.$client = new VirtueClient()
      .on("error", window.$log.error)
      .on("warn", window.$log.warn)
      .on("info", window.$log.info)
      .on("debug", window.$log.debug);

    Vue.prototype.$client = window.$client;
  },
};

export default plugin;
