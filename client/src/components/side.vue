<template>
  <aside class="virtue-menu">
    <div class="tabs-container">
      <ul>
        <li :class="{ active: tab === 'settings' }" @click.stop.prevent="change('settings')">
          <i class="fas fa-sliders-h" />
          <span>{{ $t('side.settings') }}</span>
        </li>
      </ul>
    </div>
    <div class="page-container">
      <virtue-settings v-if="tab === 'settings'" />
    </div>
  </aside>
</template>

<style lang="scss">
  .virtue-menu {
    width: $side-width;
    background-color: $background-primary;
    flex-shrink: 0;
    max-height: 100%;
    max-width: 100%;
    display: flex;
    flex-direction: column;

    .tabs-container {
      background: $background-tertiary;
      height: $menu-height;
      max-height: 100%;
      max-width: 100%;
      display: flex;
      flex-shrink: 0;

      ul {
        display: inline-block;
        padding: 16px 0 0 0;

        li {
          background: $background-secondary;
          border-radius: 3px 3px 0 0;
          border-bottom: none;
          display: inline-block;
          padding: 5px 10px;
          margin-right: 4px;
          font-weight: 600;
          cursor: pointer;

          i {
            margin-right: 4px;
            font-size: 10px;
          }

          &.active {
            background: $background-primary;
          }
        }
      }
    }

    .page-container {
      max-height: 100%;
      flex-grow: 1;
      display: flex;
      overflow: auto;
      padding-top: 5px;
    }
  }
</style>

<script lang="ts">
  import { Vue, Component } from 'vue-property-decorator'

  import Settings from '~/components/settings.vue'

  @Component({
    name: 'virtue',
    components: {
      'virtue-settings': Settings,
    },
  })
  export default class extends Vue {
    get tab() {
      return this.$accessor.client.tab
    }

    change(tab: string) {
      this.$accessor.client.setTab(tab)
    }
  }
</script>
