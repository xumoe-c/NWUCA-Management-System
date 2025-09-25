import axios from 'axios';

const http = axios.create({
    baseURL: '/api/v1',
    timeout: 10000,
});

http.interceptors.request.use((config) => {
    const token = localStorage.getItem('token');
    if (token) {
        config.headers = config.headers || {};
        config.headers['Authorization'] = `Bearer ${token}`;
    }
    return config;
});

http.interceptors.response.use(
    (res) => res,
    (err) => {
        if (err?.response?.status === 401) {
            localStorage.removeItem('token');
            // 在 H5 下可以直接跳转登录页
            if (location.pathname !== '/login') {
                const redirect = encodeURIComponent(location.pathname + location.search);
                location.href = `/login?redirect=${redirect}`;
            }
        }
        return Promise.reject(err);
    },
);

export default http;
