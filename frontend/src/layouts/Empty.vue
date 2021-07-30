<template>
  <div class="main-wrapper">
    <base-header/>
    <div class="content">
      <router-view/>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import BaseHeader from '@/components/BaseHeader.vue'

export default Vue.extend({
  name: 'Empty',
  components: { BaseHeader }
})
</script>

<style scoped>
.main-wrapper {
  min-height: 100vh;
  background: #eee;
}
.content {
  margin: 0 auto;
  max-width: 80vw;
}
</style>
