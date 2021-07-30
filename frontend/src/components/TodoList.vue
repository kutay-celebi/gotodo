<template>
  <div class="todo-list-wrapper">
    <button class="btn success">
      <span><i class="ri-add-line"/>New</span>
    </button>
    <ul v-if="todos && todos.length > 0">
      <li v-for="(todo,index) in todos" :key="`todo-${index}`">
        <div>
          <div class="todo-title">{{ todo.title }}</div>
          <div class="todo-description">{{ todo.description }}</div>
        </div>
        <div class="actions">
          <button class="btn success">
            <span><i class="ri-check-line"/> Complete</span>
          </button>
          <button class="btn danger">
            <span><i class="ri-delete-bin-2-line"/> Delete</span>
          </button>
        </div>
      </li>
    </ul>
    <div v-else class="empty-message">
      No Record
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropType } from 'vue'
import { Todo } from '@/model/Todo'

export default Vue.extend({
  name: 'TodoList',
  model: {
    prop: 'todos'
  },
  props: {
    todos: {
      type: Array as PropType<Todo[]>
    }
  }
})
</script>

<style scoped>
.todo-list-wrapper {
  padding: 1rem;
  border: 1px solid #cecece;
}

.todo-list-wrapper ul {
  list-style: none;
  margin: 0;
  padding: 0;
}

.todo-list-wrapper ul li {
  display: flex;
  line-height: 2rem;
  border: 1px solid #efefef;
  margin: .25rem 0;
  padding: 0 1rem;
}

.todo-list-wrapper ul li .todo-title {
  font-weight: 600;
  font-size: 18px;
  padding: .5rem 0;
}

.todo-list-wrapper ul li .todo-description {
  font-size: 14px;
  color: #797878;
}

.actions {
  display: flex;
  align-items: center;
  justify-content: center;
  width: 40%;
  text-align: center;
  flex-wrap: wrap;
}

.empty-message {
  font-weight: 600;
  text-align: center;
}
</style>
