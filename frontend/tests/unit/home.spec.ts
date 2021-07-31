import { mount } from '@vue/test-utils'
import Home from '@/views/Home.vue'
import MockAdapter from 'axios-mock-adapter'
import service from '@/util/api'
import flushPromises from 'flush-promises'

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

describe('Home.vue', () => {
  const mock = new MockAdapter(service)

  mock.onGet('/api/todo/list').reply(200, mockTodos)

  it('renders a child component via routing', async () => {
    const wrapper = mount(Home)
    await flushPromises()
    expect(wrapper.vm.$data.todoList).toEqual(mockTodos)
  })
})
