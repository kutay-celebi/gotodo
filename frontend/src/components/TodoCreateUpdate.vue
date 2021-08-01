<template>
  <div class="create-update-wrapper">
    <button id="show-container-btn" class="btn success" @click="showContainer = !showContainer">
      <span><i class="ri-add-line"/>New</span>
    </button>
    <transition name="fade-transform">
      <div v-if="showContainer">

        <transition name="fade-transform">
          <div v-if="errors && errors.length > 0">
            <div class="error-container">
              <i class="ri-error-warning-line"/>
              <div class="messages">
                <div v-for="(error, index) in errors" :key="`e-${index}`">{{ error }}</div>
              </div>
            </div>
          </div>
        </transition>

        <form id="todo-form" class="todo-form" @submit.prevent="handleFormSubmit">
          <div class="form-field">
            <input id="todo-title" class="text-input" placeholder="Title" v-model="todoModel.title">
          </div>
          <div class="form-field">
            <textarea id="todo-description" class="text-input" placeholder="Description" v-model="todoModel.description"/>
          </div>

          <button id="submit-button" class="btn success" type="submit">
            <span><i class="ri-save-line"></i>Save Todo</span>
          </button>
        </form>
      </div>
    </transition>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { Todo } from '@/model/Todo'

export default Vue.extend({
  name: 'TodoCreateUpdate',
  data () {
    return {
      showContainer: false,
      errors: [] as string[],
      todoModel: {} as Todo
    }
  },
  methods: {
    async handleFormSubmit () {
      if ((!this.todoModel.title || this.todoModel.title.trim() === '') && !this.errors.some(error => error === 'Please fill the title field.')) {
        this.errors.push('Please fill the title field.')
        return
      } else if (this.todoModel.title && this.todoModel.title.trim() !== '') {
        this.errors = []
      }

      try {
        await this.$emit('saveTodo', this.todoModel)
      } finally {
        this.todoModel = {}
        this.showContainer = false
      }
    }
  }
})
</script>

<style scoped>
.create-update-wrapper {
}

.todo-form {
  margin: 1rem 0;
}

.error-container {
  border: 1px solid #FCA5A5;
  border-radius: .25rem;
  padding: 1rem;
  margin: 1rem 0;
  display: flex;
  align-items: center;

}

.error-container i {
  color: #FCA5A5;
  font-size: 1.5rem;
  margin: 0 1rem
}
</style>
