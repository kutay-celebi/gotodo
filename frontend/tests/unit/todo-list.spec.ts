import { shallowMount } from '@vue/test-utils'
import TodoList from '@/components/TodoList.vue'
import { Todo } from '@/model/Todo'

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


})
