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

const syncVuetifyTheme = () => {
  if (!vuetify?.framework?.theme) return
  const mediaQuery = window.matchMedia("(prefers-color-scheme: dark)")
  const updateTheme = (event) => {
    const isDark = event.matches
    vuetify.framework.theme.dark = isDark
    document.documentElement.classList.toggle("theme--dark", isDark)
  }
  updateTheme(mediaQuery)
  if (mediaQuery.addEventListener) {
    mediaQuery.addEventListener("change", updateTheme)
  } else if (mediaQuery.addListener) {
    mediaQuery.addListener(updateTheme)
  }
}

initializeGTMConsent()
syncVuetifyTheme()

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

new Vue({
  router,
  store,
  vuetify,
  render: (h) => h(App),
}).$mount("#app")
