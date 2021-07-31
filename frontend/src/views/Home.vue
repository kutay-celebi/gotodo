<template>
  <div class="home py-8">
    <card>
      <span slot="title">
        Todo List
      </span>
      <todo-create-update/>
      <todo-list :todos="todoList" @todoComplete="completeTodo"/>
    </card>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import Card from '@/components/Card.vue'
import TodoList from '@/components/TodoList.vue'
import { Todo } from '@/model/Todo'
import TodoCreateUpdate from '@/components/TodoCreateUpdate.vue'
import service from '@/util/api'

export default Vue.extend({
  name: 'Home',
  components: { TodoCreateUpdate, TodoList, Card },
  data () {
    return {
      todoList: [] as Todo[]
    }
  },
  async beforeMount () {
    await this.getTodoList()
  },
  methods: {
    async getTodoList () {
      await service.get('/api/todo/list')
        .then(response => {
          this.todoList = response.data
        })
    }
  }
})
</script>
