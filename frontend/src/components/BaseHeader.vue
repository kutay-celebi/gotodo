<template>
  <!-- This example requires Tailwind CSS v2.0+ -->
  <div class="header">
    <div class="nav-container">
      <div class="nav-title">
        <div>Todo Applciation</div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'BaseHeader'
})
</script>

<style scoped>
.header {
  position: relative;
  background: #2c3e50;
}

.nav-title {
  display: flex;
  justify-content: start;
  flex: 1;
}

.nav-container {
  display: flex;
  justify-content: space-between;
  align-items: center;
  border-bottom: 2px solid #ddd;
  padding: 1.5rem 2rem;
  color: #ddd;
}

.nav-item {
  color: #ddd;
}
</style>
