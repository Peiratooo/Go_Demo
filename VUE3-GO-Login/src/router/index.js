import { createRouter, createWebHistory } from 'vue-router'

const routes = [
    {
        path: "/",
        name: "home",
        component: () => import('../components/Home.vue'),
    },
    {
        path: "/account",
        name: "account",
        component: () => import('../components/Account.vue'),
        children: [
            {
                path: "/login",
                name: "login",
                component: () => import('../components/Login.vue'),
            },
            {
                path: "/regist",
                name: "regist",
                component: () => import('../components/Regist.vue'),
            }
        ]
    }
]

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes,
})

const whiteList = ['/account', '/login', '/regist']

router.beforeEach((to, from, next) => {
    // console.log(to.path);
    // console.log(localStorage.getItem('token'));
    if (localStorage.getItem('Token')) {
        console.log(localStorage.getItem('Token'));
        next()
    } else {
        if (whiteList.includes(to.path)) {
            next()
        } else {
            next('/login')
        }
    }
})

export default router