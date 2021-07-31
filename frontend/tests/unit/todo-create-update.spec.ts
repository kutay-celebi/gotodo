import { shallowMount } from '@vue/test-utils'
import TodoCreateUpdate from '@/components/TodoCreateUpdate.vue'
import Vue from 'vue'

describe('TodoCreateUpdate.vue', () => {
  it('render', () => {
    const wrapper = shallowMount(TodoCreateUpdate)
    expect(wrapper.find('.create-update-wrapper').exists()).toBeTruthy()
    expect(wrapper.vm.$data.showContainer).toBeFalsy()
  })

  it('show container', () => {
    const wrapper = shallowMount(TodoCreateUpdate)
    wrapper.find('#show-container-btn').trigger('click')
    expect(wrapper.vm.$data.showContainer).toBeTruthy()
  })

  it('fill form field', async () => {
    const wrapper = shallowMount(TodoCreateUpdate, {
      data () {
        return {
          showContainer: true
        }
      }
    })

    const title = 'test title'
    const description = 'test description'

    // set form fields
    wrapper.find('#todo-title').setValue(title)
    wrapper.find('#todo-description').setValue(description)
    await Vue.nextTick()

    expect(wrapper.vm.$data.todoModel.title).toEqual(title)
    expect(wrapper.vm.$data.todoModel.description).toEqual(description)
  })

  it('validate if title field is not empty', async () => {
    const wrapper = shallowMount(TodoCreateUpdate, {
      data () {
        return {
          showContainer: true
        }
      }
    })

    const title = ''

    // set form fields
    wrapper.find('#todo-title').setValue(title)
    wrapper.find('#todo-form').trigger('submit.prevent')
    await Vue.nextTick()

    expect(wrapper.find('.error-container').exists()).toBeTruthy()
  })

  it('should validation message should not be multiplexed', async () => {
    const wrapper = shallowMount(TodoCreateUpdate, {
      data () {
        return {
          showContainer: true
        }
      }
    })

    const title = ''

    // set form fields
    wrapper.find('#todo-title').setValue(title)
    wrapper.find('#todo-form').trigger('submit.prevent')
    await Vue.nextTick()
    wrapper.find('#todo-form').trigger('submit.prevent')
    await Vue.nextTick()

    expect(wrapper.vm.$data.errors.length).toEqual(1)
  })
})
