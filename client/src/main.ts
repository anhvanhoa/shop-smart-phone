import './assets/main.css'
import { createApp } from 'vue'
import App from './App.vue'
import VueTippy from 'vue-tippy'
import router from './router'

const app = createApp(App)

app.use(router)

app.use(
    VueTippy,
    {
        directive: 'tippy', // => v-tippy
        componentSingleton: 'tippy-singleton', // => <tippy-singleton/>,
        defaultProps: {
            placement: 'auto-end',
            allowHTML: true,
        }
    }
)

app.mount('#app')
