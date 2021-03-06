<template>
  <div id="todo-list-comp">
    <todo-create-update @saveTodo="saveTodo"/>
    <div class="todo-list-wrapper">
      <div v-if="todos && todos.length > 0">
        <transition-group id="todo-list" name="fade-transform" tag="ul" appear>
          <li v-for="(todo,index) in todos" :key="`todo-${index}`">
            <div class="todo-check-container" v-if="todo.completed">
              <i class="ri-check-line"/>
            </div>
            <div class="todo-container">
              <div class="todo-title">{{ todo.title }}</div>
              <div class="todo-description">{{ todo.description }}</div>
            </div>
            <div class="actions">
              <button v-if="!todo.completed" id="complete-btn" class="btn success" @click="completeTodo(todo)" :disabled="loading">
                <span><i class="ri-check-line"/> Complete</span>
              </button>
            </div>
          </li>
        </transition-group>
      </div>
      <div v-else class="empty-message">
        No Record
      </div>
      <transition name="fade" mode="out-in">
        <span v-show="loading" class="loading-overlay">
          <i class="ri-loader-2-line spin"></i>
          <div>Loading...</div>
        </span>
      </transition>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Todo } from '@/model/Todo'
import service from '@/util/api'
import TodoCreateUpdate from '@/components/TodoCreateUpdate.vue'

export default Vue.extend({
  name: 'TodoList',
  components: { TodoCreateUpdate },
  data () {
    return {
      loading: false,
      todos: [] as Todo[]
    }
  },
  async beforeMount () {
    await this.getTodoList()
  },
  methods: {
    async getTodoList () {
      await service.get('/api/todo/list')
        .then(response => {
          this.todos = response.data
        })
    },
    async completeTodo (todo: Todo) {
      this.loading = true
      try {
        await service.get(`/api/todo/complete/${todo.id}`)
          .then(async () => {
            await this.$toast.success(`${todo.title} has completed`)
            await this.getTodoList()
          })
      } finally {
        this.loading = false
      }
    },
    async saveTodo (todo: Todo) {
      this.loading = true
      try {
        await service.post('/api/todo/create', todo)
          .then(async () => {
            await this.$toast.success(`${todo.title} successfully saved.`)
            await this.getTodoList()
          })
      } finally {
        this.loading = false
      }
    }
  }

})
</script>

<style scoped>
.todo-list-wrapper {
  padding: 1rem;
  border: 1px solid #cecece;
}

.todo-container {
  width: 100%;
}

.todo-check-container {
  display: flex;
  align-items: center;
  font-size: 24px;
  margin-right: 1rem;
  color: #27ae60;
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
  float: right;
}

.empty-message {
  font-weight: 600;
  text-align: center;
}
</style>
