/**
 * plugins/vuetify.ts
 *
 * Framework documentation: https://vuetifyjs.com`
 */

// Styles
import '@mdi/font/css/materialdesignicons.css'
import 'vuetify/styles'

// Composables
import {createVuetify, ThemeDefinition} from 'vuetify'
import * as components from 'vuetify/components'
import * as directives from 'vuetify/directives'
import {md3} from 'vuetify/blueprints'
import {VSkeletonLoader} from "vuetify/labs/VSkeletonLoader";
// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides

const myCustomLightTheme: ThemeDefinition = {
  dark: false,
  colors: {
    background: '#FFF3E0',
    surface: '#FFF3E0',
    primary: '#1D5D9B',
    secondary: '#75C2F6',
    error: '#FFF3E0',
    info: '#FFF3E0',
    success: '#FFF3E0',
    warning: '#FFF3E0',
  },
}
export default createVuetify({
  components: {...components, VSkeletonLoader},
  directives,
  blueprint: md3,
  theme: {
    defaultTheme: 'myCustomLightTheme',
    themes: {
      myCustomLightTheme,
    },
  },
})

