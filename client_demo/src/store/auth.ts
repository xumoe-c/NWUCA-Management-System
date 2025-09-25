import { defineStore } from 'pinia';

export const useAuthStore = defineStore('auth', {
    state: () => ({
        token: localStorage.getItem('token') || '',
        role: localStorage.getItem('role') || 'member',
        email: '',
    }),
    getters: {
        isAuthed: (s) => !!s.token,
        isAdmin: (s) => s.role === 'admin',
    },
    actions: {
        setAuth(token: string, role: string) {
            this.token = token;
            this.role = role || 'member';
            localStorage.setItem('token', token);
            localStorage.setItem('role', this.role);
        },
        clear() {
            this.token = '';
            this.role = 'member';
            localStorage.removeItem('token');
            localStorage.removeItem('role');
        },
    },
});
