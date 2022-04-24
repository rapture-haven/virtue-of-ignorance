<template>
  <div id="virtue">
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
#app {
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
//
@Component({
  name: "virtue",
  components: {
    "virtue-connect": Connect,
  },
})
export default class extends Vue {
  get connected() {
    return this.$accessor.connected;
  }
}
</script>
