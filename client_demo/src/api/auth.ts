import http from './http';

export interface RegisterPayload {
    username: string;
    email: string;
    password: string;
}

export interface ApiResponse<T> {
    code: number;
    msg: string;
    data: T;
}

export async function registerApi(payload: RegisterPayload) {
    // baseURL 已为 '/api/v1'，此处只需相对路径，避免出现 '/api/v1/api/v1/register'
    const res = await http.post<ApiResponse<{ user_id: number }>>('/register', payload);
    return res.data;
}

export async function loginApi(payload: { email: string; password: string }) {
    // 同理，仅使用相对路径
    const res = await http.post<ApiResponse<{ token: string }>>('/login', payload);
    return res.data;
}
