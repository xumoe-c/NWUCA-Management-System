// userStore.js
import { defineStore } from 'pinia';

export const useUserStore = defineStore('user', {
    state: () => ({
        token: null,
        userInfo: {},
    }),
    actions: {
        setToken(token) {
            this.token = token;
        },
        logout() {
            this.token = null;
            this.userInfo = {};
        },
    },
});
