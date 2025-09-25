import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router';
import Login from '../views/Login.vue';
import Register from '../views/Register.vue';
import Dashboard from '../views/Dashboard.vue';
import Departments from '../views/Departments.vue';
import Positions from '../views/Positions.vue';
import Members from '../views/Members.vue';
import Assignments from '../views/Assignments.vue';
import Profile from '../views/Profile.vue';

const routes: RouteRecordRaw[] = [
    { path: '/', redirect: '/dashboard' },
    { path: '/login', component: Login },
    { path: '/register', component: Register },
    { path: '/dashboard', component: Dashboard, meta: { auth: true } },
    { path: '/departments', component: Departments, meta: { auth: true } },
    { path: '/positions', component: Positions, meta: { auth: true } },
    { path: '/members', component: Members, meta: { auth: true } },
    { path: '/assignments', component: Assignments, meta: { auth: true } },
    { path: '/profile', component: Profile, meta: { auth: true } },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, _from, next) => {
    const token = localStorage.getItem('token');
    if (to.meta.auth && !token) {
        next({ path: '/login', query: { redirect: to.fullPath } });
    } else {
        next();
    }
});

export default router;
