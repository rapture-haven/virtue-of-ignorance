<template>
  <div id="virtue">
    <main class="virtue-main">
      <div v-if="!hideControls" class="header-container">
          <virtue-header />
      </div>
    </main>
    <virtue-side v-if="!hideControls && side" />
    <virtue-connect v-if="!connected" />
    <notifications
      group="virtue"
      position="top left"
      style="top: 50px; pointer-events: none"
      :ignoreDuplicates="true"
    />
  </div>
</template>

<style lang="scss">
#virtue {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}

#nav {
  padding: 30px;

  a {
    font-weight: bold;
    color: #2c3e50;

    &.router-link-exact-active {
      color: #42b983;
    }
  }
}
</style>

<script lang="ts">
import { Vue, Component, Ref, Watch } from "vue-property-decorator";

import Connect from "@/components/connect.vue";
import Header from '@/components/header.vue'
import Settings from '~/components/settings.vue'

@Component({
  name: "virtue",
  components: {
    "virtue-connect": Connect,
    'virtue-header': Header,
    'virtue-settings': Settings,
  },
})
export default class extends Vue {
  get hideControls() {
      return !!new URL(location.href).searchParams.get('cast')
  }

  @Watch('hideControls', { immediate: true })
  onHideControls(enabled: boolean) {
    if (enabled) {
      // this.$accessor.video.setMuted(false)
      // this.$accessor.settings.setSound(false)
    }
  }

  get about() {
      return this.$accessor.client.about
  }

  get side() {
      return this.$accessor.client.side
    }

  get connected() {
    return this.$accessor.connected;
  }
}
</script>
