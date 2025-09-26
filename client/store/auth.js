export function setToken(token) {
  try { uni.setStorageSync('auth_token', token) } catch (e) {}
}

export function getToken() {
  try { return uni.getStorageSync('auth_token') || '' } catch (e) { return '' }
}

export function clearToken() {
  try { uni.removeStorageSync('auth_token') } catch (e) {}
}

export function setUser(user) {
  try { uni.setStorageSync('auth_user', user) } catch (e) {}
}

export function getUser() {
  try { return uni.getStorageSync('auth_user') || null } catch (e) { return null }
}

export function clearUser() {
  try { uni.removeStorageSync('auth_user') } catch (e) {}
}

export function isLoggedIn() {
  return !!getToken()
}
