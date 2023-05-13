import { HOST_URL, authApi, publicApi } from '../api'

class ProductService {
  static async getPopularBooks() {
    const res = await publicApi.post('/popular')
    return res.data
  }

  static async getDiscountedBooks() {
    const res = await publicApi.post('/discount')
    return res.data
  }

  static async addFavourite(options) {
    const res = await authApi.post('/addFavourite', options)
    return res.data
  }

  static async removeFavourite(options) {
    const res = await authApi.post('/deleteFavourite', options)
    return res.data
  }

  static async getFavourite(pageSize = 16, page = 'page=1') {
    const res = await authApi.get(`/getFavourites?&page_size=${pageSize}&${page ? page :'page=1'}`)
    return res.data
  }

  static async getCart(ids) {
    const res = await publicApi.post(`/getCart`, ids)
    return res.data
  }

  static async getSearch(query = "page=1", pageSize = 9) {
    const res = await publicApi.get(
      `/search?&${query}&page_size=${pageSize}`
    )
    return res.data
  }

  static async getBook(id) {
    const res = await publicApi.get(`/book/${id}`)
    return res.data
  }

  static async createOrder(adress, ids) {
    let book_ids = []
    let amounts = []
    for (let i = 0; i < ids.length; i++) {
      book_ids.push(ids[i].id)
      amounts.push(ids[i].amount)
    }
    let data = {Adress: adress.Adress, Book_IDs: book_ids, Amounts: amounts} 
    const res = await authApi.post('/createOrder', data)
    return res.data
  }

  static async getOrders() {
    const res = await authApi.get('/getOrders')
    return res.data
  }

  static async getFilters() {
    const res = await publicApi.post('/filters', )
    return res.data
  }

  //static async deleteOrder(id) {
  //  const res = await publicApi.post('/deleteOrder', id)
  //  return res.data
  // }
}

export default ProductService
