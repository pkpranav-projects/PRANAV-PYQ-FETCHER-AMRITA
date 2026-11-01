import DefaultTheme from 'vitepress/theme'
import './style.css'
import DynamicBackground from './DynamicBackground.vue'
import { h } from 'vue'

export default {
  extends: DefaultTheme,
  Layout() {
    return h(DefaultTheme.Layout, null, {
      'layout-top': () => h(DynamicBackground),
    })
  },
}
