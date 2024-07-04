import 'normalize.css/normalize.css'
import 'element-plus/dist/index.css'
import './elementplus.css'
import { createApp } from 'vue'
import { createPinia } from 'pinia'
import ElementPlus from 'element-plus'
import App from './App.vue'
import 'virtual:svg-icons-register'

const app = createApp(App)
const pinia = createPinia()

app.use(ElementPlus)
app.use(pinia)
app.mount('#app')
