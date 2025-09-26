import { http } from './http'

export function loginApi(payload) {
  // Expected payload: { username, password }
  return http.post('/api/v1/login', payload)
}

export function getMeApi() {
  return http.get('/api/v1/users/me')
}
