// Simple uni.request wrapper with token injection and 401 handling
function getBaseUrl() {
  try {
    const saved = uni.getStorageSync('API_BASE');
    if (saved) return saved;
  } catch (e) {}
  // Allow injecting a global at runtime for H5
  if (typeof window !== 'undefined' && window.__API_BASE__) return window.__API_BASE__;
  return 'http://localhost:8080';
}

function getToken() {
  try {
    return uni.getStorageSync('auth_token') || '';
  } catch (e) {
    return '';
  }
}

function clearAuthAndRedirect() {
  try { uni.removeStorageSync('auth_token'); uni.removeStorageSync('auth_user'); } catch (e) {}
  if (typeof uni !== 'undefined') {
    // Avoid redirect loop if already on login
    const pages = getCurrentPages();
    const current = pages && pages.length ? pages[pages.length - 1] : null;
    const route = current && (current.route || (current.$page && current.$page.path));
    if (route !== 'pages/login/index') {
      uni.reLaunch({ url: '/pages/login/index' });
    }
  }
}

export function request({ url, method = 'GET', data = {}, header = {}, timeout = 20000 }) {
  return new Promise((resolve, reject) => {
    const token = getToken();
    const headers = Object.assign({ 'Content-Type': 'application/json' }, header);
    if (token) headers['Authorization'] = `Bearer ${token}`;

    uni.request({
      url: url.startsWith('http') ? url : `${getBaseUrl()}${url}`,
      method,
      data,
      header: headers,
      timeout,
      success: (res) => {
        // Prefer backend unified envelope: { code, message, data }
        if (res.statusCode === 401 || (res.data && (res.data.code === 401 || res.data.code === '401'))) {
          uni.showToast({ title: '登录已过期，请重新登录', icon: 'none' });
          clearAuthAndRedirect();
          return reject(res);
        }
        resolve(res.data !== undefined ? res.data : res);
      },
      fail: (err) => {
        uni.showToast({ title: '网络错误，请稍后重试', icon: 'none' });
        reject(err);
      }
    });
  });
}

export const http = {
  get: (url, params = {}, header = {}) => request({ url, method: 'GET', data: params, header }),
  post: (url, data = {}, header = {}) => request({ url, method: 'POST', data, header }),
  put: (url, data = {}, header = {}) => request({ url, method: 'PUT', data, header }),
  delete: (url, data = {}, header = {}) => request({ url, method: 'DELETE', data, header }),
};

export default http;
