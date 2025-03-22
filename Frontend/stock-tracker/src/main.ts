import { createApp } from 'vue'
import { createPinia } from 'pinia'

import App from './App.vue'
import router from "./router"; // Importamos Vue Router
import "./assets/tailwind.css";


const app = createApp(App)

app.use(createPinia())
app.use(router)

app.mount('#app')
