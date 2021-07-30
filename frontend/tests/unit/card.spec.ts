import { shallowMount } from '@vue/test-utils'
import Card from '@/components/Card.vue'

describe('Card.vue', () => {
  it('render', () => {
    const wrapper = shallowMount(Card)
    expect(wrapper.find('.card').exists()).toBeTruthy()
  })

  it('slots', () => {
    const slotWrapper = shallowMount(Card, {
      slots: {
        title: 'Test'
      }
    })

    const actual = slotWrapper.find('.title').text()
    expect(actual).toEqual('Test')
  })
})
