<template>
  <div class="home py-8">
    <card>
      <span slot="title">
        Todo List
      </span>
      <todo-list/>
    </card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Card from '@/components/Card.vue'
import TodoList from '@/components/TodoList.vue'

export default Vue.extend({
  name: 'Home',
  components: { TodoList, Card }
})
</script>
