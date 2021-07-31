<template>
  <div class="create-update-wrapper">
    <button id="show-container-btn" class="btn success" @click="showContainer = !showContainer">
      <span><i class="ri-add-line"/>New</span>
    </button>
    <transition name="fade-transform">
      <div v-if="showContainer">
        <form  id="todo-form" class="todo-form" @submit.prevent="handleFormSubmit">
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
      todoModel: {} as Todo
    }
  },
  methods: {
    handleFormSubmit () {
      console.log(this.todoModel)
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
</style>
