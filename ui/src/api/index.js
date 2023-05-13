import axios from 'axios'

const HOST_URL = 'http://localhost:83'
const API_URL = `${HOST_URL}/api`

const publicApi = axios.create({
  baseURL: API_URL,
})

const authApi = axios.create({
  withCredentials: true,
  baseURL: API_URL,
})

//authApi.interceptors.request.use((config) => {
  //config.headers.Authorization = `JWT ${localStorage.getItem('token')}`
  //return config
//})

authApi.interceptors.response.use(
  (config) => {
    return config
  },
  async (error) => {
    const originalReq = error.config
    console.log(originalReq)
    if (
      error.response.status === 401 &&
      error.config &&
      !error.config._isRetry
    ) {
      originalReq._isRetry = true
      try {
        const res = await authApi.post('/refresh')
        return authApi.request(originalReq)
      } catch (error) {
        throw error
      }
    }
    throw error
  }
)

export { HOST_URL, API_URL, publicApi, authApi }
