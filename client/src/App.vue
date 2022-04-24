<template>
  <div id="virtue">
    <main class="virtue-main">
      <div v-if="!hideControls" class="header-container">
          <virtue-header />
      </div>
      <div v-if="!hideControls" class="room-container">
          <virtue-members />
          <div class="room-menu">
            <div class="settings">
              <virtue-menu />
            </div>
            <div class="controls">
              <virtue-controls :shakeKbd="shakeKbd" />
            </div>
          </div>
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
    position: absolute;
    top: 0;
    left: 0;
    right: 0;
    bottom: 0;
    max-width: 100vw;
    max-height: 100vh;
    flex-direction: row;
    display: flex;

    .virtue-main {
      min-width: 360px;
      max-width: 100%;
      flex-grow: 1;
      flex-direction: column;
      display: flex;
      overflow: auto;

      .header-container {
        background: $background-tertiary;
        height: $menu-height;
        flex-shrink: 0;
        display: flex;
      }

      .video-container {
        background: rgba($color: #000, $alpha: 0.4);
        max-width: 100%;
        flex-grow: 1;
        display: flex;
      }

      .room-container {
        background: $background-tertiary;
        height: $controls-height;
        max-width: 100%;
        flex-shrink: 0;
        flex-direction: column;
        display: flex;

        .room-menu {
          max-width: 100%;
          flex: 1;
          display: flex;

          .settings {
            margin-left: 10px;
            flex: 1;
            justify-content: flex-start;
            align-items: center;
            display: flex;
          }

          .controls {
            flex: 1;
            justify-content: center;
            align-items: center;
            display: flex;
          }

          .emotes {
            margin-right: 10px;
            flex: 1;
            justify-content: flex-end;
            align-items: center;
            display: flex;
          }
        }
      }
    }
  }

  @media only screen and (max-width: 600px) {
    #virtue.expanded {
      .virtue-main {
        transform: translateX(calc(-100% + 65px));

        video {
          display: none;
        }
      }

      .virtue-menu {
        position: absolute;
        top: 0;
        right: 0;
        bottom: 0;
        left: 65px;
        width: calc(100% - 65px);
      }
    }
  }

  @media only screen and (max-width: 768px) {
    #virtue .virtue-main .room-container {
      display: none;
    }
  }
</style>

<script lang="ts">
import { Vue, Component, Ref, Watch } from "vue-property-decorator";

import Connect from "@/components/connect.vue";
import Header from '@/components/header.vue'
import Side from '~/components/side.vue'
import Menu from '~/components/menu.vue'
import Controls from '~/components/controls.vue'
import Members from '~/components/members.vue'

@Component({
  name: "virtue",
  components: {
    "virtue-connect": Connect,
    'virtue-header': Header,
    'virtue-side': Side,
    'virtue-menu': Menu,
    'virtue-controls': Controls,
    'virtue-members': Members,
  },
})
export default class extends Vue {
  shakeKbd = false

  get hideControls() {
      return !!new URL(location.href).searchParams.get('cast')
  }

  @Watch('hideControls', { immediate: true })
  onHideControls(enabled: boolean) {
    if (enabled) {
      this.$accessor.video.setMuted(false)
    }
  }

  controlAttempt() {
    if (this.shakeKbd || this.$accessor.remote.hosted) return

    this.shakeKbd = true
    window.setTimeout(() => (this.shakeKbd = false), 5000)
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
