import { createApp } from "vue"
import App from "./App.vue"
import "./assets/main.css"

import { createRouter, createWebHistory } from "vue-router"
import Home from "./views/Home.vue"

const routes = [{ path: "/", component: Home }]
const router = createRouter({
  history: createWebHistory(),
  routes,
})

createApp(App).use(router).mount("#app")
