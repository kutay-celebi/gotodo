import { shallowMount } from '@vue/test-utils'
import TodoList from '@/components/TodoList.vue'
import { Todo } from '@/model/Todo'

describe('TodoList.vue', () => {
  it('render', () => {
    const wrapper = shallowMount(TodoList)
    expect(wrapper.find('.todo-list-wrapper').exists()).toBeTruthy()
  })

  it('list count', () => {
    const todoList: Todo[] = [{ title: 'test', description: 'test' }, { title: 'test2', description: 'test2' }]
    const wrapper = shallowMount(TodoList, { propsData: { todos: todoList } })
    const liArray = wrapper.findAll('li')
    expect(liArray.length).toEqual(todoList.length)
  })

  it('empty text', () => {
    const todoList: Todo[] = []
    const wrapper = shallowMount(TodoList, { propsData: { todos: todoList } })
    const liArray = wrapper.findAll('li')

    expect(liArray.length).toEqual(todoList.length)
    const actual = wrapper.find('.empty-message').element as HTMLElement
    expect(actual.textContent?.trim()).toEqual('No Record')
  })
})
