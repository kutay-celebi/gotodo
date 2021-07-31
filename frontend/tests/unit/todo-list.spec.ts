import { createLocalVue, mount, shallowMount } from '@vue/test-utils'
import TodoList from '@/components/TodoList.vue'
import { Todo } from '@/model/Todo'
import MockAdapter from 'axios-mock-adapter'
import service from '@/util/api'
import Home from '@/views/Home.vue'
import flushPromises from 'flush-promises'
import Vue from 'vue'
import VueToastificationPlugin from 'vue-toastification'

describe('TodoList.vue', () => {
  it('render', () => {
    const wrapper = shallowMount(TodoList)
    expect(wrapper.find('.todo-list-wrapper').exists()).toBeTruthy()
  })

  it('list count', () => {
    const todoList: Todo[] = [
      { title: 'test', description: 'test', completed: false },
      { title: 'test2', description: 'test2', completed: true }
    ]
    const wrapper = shallowMount(TodoList, { propsData: { todos: todoList } })
    const liArray = wrapper.findAll('li')
    expect(liArray.length).toEqual(todoList.length)
  })

  it('should empty text rendered', () => {
    const todoList: Todo[] = []
    const wrapper = shallowMount(TodoList, { propsData: { todos: todoList } })
    const liArray = wrapper.findAll('li')

    expect(liArray.length).toEqual(todoList.length)
    const actual = wrapper.find('.empty-message').element as HTMLElement
    expect(actual.textContent?.trim()).toEqual('No Record')
  })

  it('should the complete button does not appear on completed todo', () => {
    const todoList: Todo[] = [
      { title: 'test', description: 'test', completed: true }
    ]
    const wrapper = shallowMount(TodoList, { propsData: { todos: todoList } })

    // check item render
    const listItem = wrapper.find('li')
    expect(listItem.exists()).toBeTruthy()

    // check if button is not render
    const todo = wrapper.find('#complete-btn')
    expect(todo.exists()).toBeFalsy()
  })

  // it('should the complete button call todoComplete event.', async () => {
  //   const todoList: Todo[] = [
  //     { id: '1', title: 'test', description: 'test', completed: false }
  //   ]
  //
  //   // trigger complete method.
  //   const wrapper = shallowMount(TodoList, { propsData: { todos: todoList } })
  //   wrapper.find('#complete-btn').trigger('click')
  //   await wrapper.vm.$nextTick()
  //
  //   expect(wrapper.emitted().todoComplete).toBeDefined()
  // })

  it('should todo complete after click complete button', async () => {
    const mock = new MockAdapter(service)
    mock.onGet('/api/todo/complete/1').reply(200)

    const todoList: Todo[] = [
      { id: '1', title: 'test', description: 'test', completed: false }
    ]

    const localVue = createLocalVue()
    localVue.use(VueToastificationPlugin)
    const wrapper = mount(TodoList, {
      propsData: { todos: todoList },
      localVue
    })

    wrapper.find('#complete-btn').trigger('click')
    expect(wrapper.vm.$data.loading).toBeTruthy()
    await flushPromises()

    expect(wrapper.vm.$data.loading).toBeFalsy()
  })
})
