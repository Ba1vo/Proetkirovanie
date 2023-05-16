import { withCredentialsApi, authApi, publicApi } from '../api'

class AuthService {
  static async login(user) {
    const res = await authApi.post('/login', user)
    return res.data
  }

  static async register(user) {
    const res = await authApi.post('/register', user)
    return res.data
  }

  static async loginCookie(user) {
    const res = await authApi.post('/cookie', user)
    console.log(res.data)
    return res.data
  }

  static async logout() {
    await authApi.post('/logout')
  }

  /*static async checkAuth() {
    const res = await withCredentialsApi.post('/login/refresh')
    return res.data
  }*/
}

export default AuthService
