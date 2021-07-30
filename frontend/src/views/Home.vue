<template>
  <div class="home py-8">
    <card>
      <span slot="title">
        Todo List
      </span>
      <todo-list v-model="todoList"/>
    </card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Card from '@/components/Card.vue'
import TodoList from '@/components/TodoList.vue'
import { Todo } from '@/model/Todo'

export default Vue.extend({
  name: 'Home',
  components: { TodoList, Card },
  data () {
    return {
      todoList: [
        { title: 'test', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi faucibus faucibus odio a volutpat. Nunc rhoncus hendrerit tortor eget auctor. Fusce ultrices sapien risus, in ultrices est pharetra sit amet. Nam tempor tortor id ipsum vulputate laoreet ac finibus libero. Proin dapibus orci eu nisl feugiat, quis venenatis magna placerat. ' },
        { title: 'test2', description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi faucibus faucibus odio a volutpat. Nunc rhoncus hendrerit tortor eget auctor. Fusce ultrices sapien risus, in ultrices est pharetra sit amet. Nam tempor tortor id ipsum vulputate laoreet ac finibus libero. Proin dapibus orci eu nisl feugiat, quis venenatis magna placerat. ' }
      ] as Todo[]
    }
  }
})
</script>
