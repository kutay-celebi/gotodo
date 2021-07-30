import { createLocalVue, mount } from '@vue/test-utils'
import App from '@/App.vue'
import VueRouter from 'vue-router'
import { routes } from '@/router'
import Home from '@/views/Home.vue'

const localVue = createLocalVue()
localVue.use(VueRouter)

describe('App', () => {
  it('renders a child component via routing', async () => {
    const router = new VueRouter({ routes })
    const wrapper = mount(App, {
      localVue,
      router
    })

    router.push('/')
    await wrapper.vm.$nextTick()

    expect(wrapper.findComponent(Home).exists()).toBe(true)
  })
})
