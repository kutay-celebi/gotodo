import { createLocalVue, mount, shallowMount } from '@vue/test-utils'
import TodoList from '@/components/TodoList.vue'
import MockAdapter from 'axios-mock-adapter'
import service from '@/util/api'
import flushPromises from 'flush-promises'
import VueToastificationPlugin from 'vue-toastification'

const mockTodos = [
  {
    title: 'test',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi faucibus faucibus odio a volutpat. Nunc rhoncus hendrerit tortor eget auctor. Fusce ultrices sapien risus, in ultrices est pharetra sit amet. Nam tempor tortor id ipsum vulputate laoreet ac finibus libero. Proin dapibus orci eu nisl feugiat, quis venenatis magna placerat. '
  },
  {
    title: 'test2',
    description: 'Lorem ipsum dolor sit amet, consectetur adipiscing elit. Morbi faucibus faucibus odio a volutpat. Nunc rhoncus hendrerit tortor eget auctor. Fusce ultrices sapien risus, in ultrices est pharetra sit amet. Nam tempor tortor id ipsum vulputate laoreet ac finibus libero. Proin dapibus orci eu nisl feugiat, quis venenatis magna placerat. '
  }
]

describe('TodoList.vue', () => {
  const mock = new MockAdapter(service)
  const mockApi = mock.onGet('/api/todo/list')

  it('should render', () => {
    mockApi.reply(200, mockTodos)
    const wrapper = shallowMount(TodoList)
    expect(wrapper.find('.todo-list-wrapper').exists()).toBeTruthy()
  })

  it('todo list should be pulled when component is rendered', async () => {
    mockApi.reply(200, mockTodos)
    const wrapper = mount(TodoList)
    await flushPromises()
    expect(wrapper.vm.$data.todos).toEqual(mockTodos)
  })

  it('list count', async () => {
    mockApi.reply(200, mockTodos)
    const wrapper = mount(TodoList)
    await flushPromises()
    const liArray = wrapper.findAll('li')
    expect(liArray.length).toEqual(mockTodos.length)
  })

  it('should empty text rendered', () => {
    mockApi.reply(200, [])
    const wrapper = shallowMount(TodoList)
    const liArray = wrapper.findAll('li')

    expect(liArray.length).toEqual(0)
    const actual = wrapper.find('.empty-message').element as HTMLElement
    expect(actual.textContent?.trim()).toEqual('No Record')
  })

  it('the completed icon should appear on completed todo.', async () => {
    mockApi.reply(200, [{ title: 'test', description: 'test', completed: true }])
    const wrapper = shallowMount(TodoList)
    await flushPromises()

    const checkElement = wrapper.find('.todo-check-container')
    expect(checkElement.exists()).toBeTruthy()
  })

  it('should the complete button does not appear on completed todo', async () => {
    mockApi.reply(200, [{ title: 'test', description: 'test', completed: true }])
    const wrapper = shallowMount(TodoList)
    await flushPromises()

    // check item render
    const listItem = wrapper.find('li')
    await flushPromises()

    expect(listItem.exists()).toBeTruthy()

    // check if button is not render
    const todo = wrapper.find('#complete-btn')
    expect(todo.exists()).toBeFalsy()
  })

  it('should todo complete after click complete button', async () => {
    mockApi.reply(200, [{ id: '1', title: 'test', description: 'test', completed: false }])
    mock.onGet('/api/todo/complete/1').reply(200)

    const localVue = createLocalVue()
    localVue.use(VueToastificationPlugin)
    const wrapper = mount(TodoList, {
      localVue
    })
    await flushPromises()

    wrapper.find('#complete-btn').trigger('click')
    expect(wrapper.vm.$data.loading).toBeTruthy()
    await flushPromises()

    expect(wrapper.vm.$data.loading).toBeFalsy()
  })

  it('save endpoint should be called', async () => {
    mockApi.reply(200, [])
    mock.onPost('/api/todo/create').reply(200)

    const localVue = createLocalVue()
    localVue.use(VueToastificationPlugin)
    const wrapper = shallowMount(TodoList, {
      localVue
    })
    await flushPromises()

    /* eslint-disable */
    // @ts-ignore
    await wrapper.vm.saveTodo({ title: 'test', description: 'test', completed: false })

    expect(mock.history.post.length).toBe(1)
    expect(wrapper.vm.$data.loading).toBeFalsy()
  })
})
