// api.js
import axios from 'axios';

const instance = axios.create({
    baseURL: 'http://localhost:8080/api/v1', // 后端服务地址
    timeout: 5000,
});

export default instance;
