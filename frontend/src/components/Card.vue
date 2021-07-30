<template>
  <div class="card">
    <h3 class="title">
      <slot name="title"/>
    </h3>

    <div class="content">
      <slot #default/>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'Card'
})
</script>

<style scoped>
.card {
  width: 100%;
  border: 1px solid #ccc;
  padding: 0 1rem;
  background: #fdfdfd;
  border-radius: 0.25rem;
}

.card .title {
  letter-spacing: 0.125rem;
  text-align: left;
  font-weight: 600;
  border-bottom: 1px solid #ccc;
  padding: 1rem 0;
}

.card .content {
  padding: 1rem;
}
</style>
