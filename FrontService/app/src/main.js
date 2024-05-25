import { createApp } from 'vue'

import App from './App.vue'
import { router } from "./router.js";

import "bootstrap/dist/js/bootstrap.js"
import "bootstrap/dist/css/bootstrap.css"

import { library } from '@fortawesome/fontawesome-svg-core'
import { fas } from '@fortawesome/free-solid-svg-icons'
import { faTwitter, faFontAwesome } from '@fortawesome/free-brands-svg-icons'
import axios from "axios"
import VueAxios from "vue-axios"

library.add(fas, faTwitter, faFontAwesome)

createApp(App).use(VueAxios, axios).use(router).mount('#app')
