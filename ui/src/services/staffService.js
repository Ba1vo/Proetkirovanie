import { HOST_URL, publicApi } from '../api'

class StaffService {
  static async deleteBook(id) {
    const res = await publicApi.post('/emp/deleteBook', id)
    return res
  }

  static async addBook(book) {
    const res = await publicApi.post('/emp/addBook', book)
    return res
  }

  static async redactBook(book) {
    const res = await publicApi.post('/emp/redactBook', book)
    return res
  }
}

export default StaffService
