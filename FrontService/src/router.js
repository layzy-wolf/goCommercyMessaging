import Chat from "./pages/Chat.vue";
import {createRouter, createMemoryHistory, createWebHashHistory} from "vue-router";
import Index from "./pages/Index.vue";
import SignIn from "./pages/SignIn.vue";
import Register from "./pages/Register.vue";

const routes = [
    {path: "/", component: Index},
    {path: "/sign", component: SignIn},
    {path: "/register", component: Register},
    {path: "/chat", component: Chat},
    {path: "/chat/:user", component: Chat},
]

export const router = createRouter({
    history: createWebHashHistory(),
    routes
})