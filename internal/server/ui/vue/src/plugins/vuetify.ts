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
// https://vuetifyjs.com/en/introduction/why-vuetify/#feature-guides

export const myCustomLightTheme: ThemeDefinition = {
  dark: false,
  colors: {
    background: '#FFF3E0',
    surface: '#FFF3E0',
    primary: '#1D5D9B',
    secondary: '#75C2F6',
  },
}

export const myCustomDarkTheme: ThemeDefinition = {
  dark: true,
  colors: {
    background: '#2D4356',
    surface: '#2D4356',
    primary: '#4682A9',
    secondary: '#EAB2A0',
  },
}
export default createVuetify({
  components: {...components},
  directives,
  blueprint: md3,
  theme: {
    defaultTheme: 'light',
    themes: {
      'light': myCustomLightTheme,
      'dark': myCustomDarkTheme
    },
  },
})

