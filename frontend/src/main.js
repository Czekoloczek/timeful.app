import Vue from "vue"
import VueWorker from "vue-worker"
import App from "./App.vue"
import router from "./router"
import store from "./store"
import vuetify from "./plugins/vuetify"
import posthogPlugin from "./plugins/posthog"
import VueGtm from "@gtm-support/vue2-gtm"
import VueMeta from "vue-meta"
import { initializeGTMConsent, hasAnalyticsConsent } from "./utils/cookie_utils"
import "./index.css"

initializeGTMConsent()

// Posthog
Vue.use(posthogPlugin)

// Google Analytics
Vue.use(VueGtm, {
  id: "GTM-M677X6V",
  vueRouter: router,
  enabled: hasAnalyticsConsent(),
})

// Site Metadata
Vue.use(VueMeta)

// Workers
Vue.use(VueWorker)

Vue.config.productionTip = false

const app = new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
})

export const syncVuetifyTheme = (manualPreference = null, vuetifyApp = app) => {
  const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)")
  const applyTheme = (isDark) => {
    vuetifyApp.$vuetify.theme.dark = isDark
    document.documentElement.classList.toggle("theme--dark", isDark)
  }
  const updateTheme = (event) => {
    applyTheme(event.matches)
  }
  if (manualPreference === "light") {
    applyTheme(false)
    return
  }
  if (manualPreference === "dark") {
    applyTheme(true)
    return
  }
  applyTheme(mediaQuery.matches)
  if (mediaQuery.addEventListener) {
    mediaQuery.addEventListener("change", updateTheme)
  } else if (mediaQuery.addListener) {
    mediaQuery.addListener(updateTheme)
  }
}

const storedTheme = localStorage.getItem("themePreference")
syncVuetifyTheme(storedTheme, app)
app.$mount("#app")
