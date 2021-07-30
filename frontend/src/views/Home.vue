<template>
  <div class="home py-8">
    <card>
      <span slot="title">
        Todo List
      </span>
      Card content
    </card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Card from '@/components/Card.vue'

export default Vue.extend({
  name: 'Home',
  components: { Card },
  data () {
    return {
      enabled: false
    }
  }
})
</script>
