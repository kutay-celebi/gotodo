import { shallowMount } from '@vue/test-utils'
import Card from '@/components/Card.vue'
import BaseHeader from '@/components/BaseHeader.vue'

describe('BaseHeader.vue', () => {
  it('render', () => {
    const wrapper = shallowMount(BaseHeader)
    expect(wrapper.find('.header').exists()).toBeTruthy()
  })
})
