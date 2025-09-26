import App from './App'
import { isLoggedIn } from './store/auth'

function isPublicPath(url) {
  return url && (url.indexOf('/pages/login/index') === 0)
}

function guard(url) {
  if (isPublicPath(url)) return true
  if (!isLoggedIn()) {
    uni.reLaunch({ url: '/pages/login/index' })
    return false
  }
  return true
}

function installInterceptors() {
  // navigateTo
  uni.addInterceptor('navigateTo', { invoke: (args) => guard(args.url) })
  // redirectTo
  uni.addInterceptor('redirectTo', { invoke: (args) => guard(args.url) })
  // reLaunch
  uni.addInterceptor('reLaunch', { invoke: (args) => guard(args.url) })
  // switchTab
  uni.addInterceptor('switchTab', { invoke: (args) => guard(args.url) })
}

// #ifndef VUE3
import Vue from 'vue'
import './uni.promisify.adaptor'
Vue.config.productionTip = false
App.mpType = 'app'
installInterceptors()
const app = new Vue({
  ...App
})
app.$mount()
// #endif

// #ifdef VUE3
import { createSSRApp } from 'vue'
export function createApp() {
  const app = createSSRApp(App)
  installInterceptors()
  return {
    app
  }
}
// #endif