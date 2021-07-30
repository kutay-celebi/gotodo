import { shallowMount } from '@vue/test-utils'
import Empty from '@/layouts/Empty.vue'

describe('Empty.vue', () => {
  it('render', () => {
    const wrapper = shallowMount(Empty, { stubs: ['router-link', 'router-view'] })
    expect(wrapper.find('.main-wrapper').exists()).toBeTruthy()
  })
})
