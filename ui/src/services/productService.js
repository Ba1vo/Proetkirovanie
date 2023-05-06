import { HOST_URL, publicApi } from '../api'

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
    const res = await publicApi.post('/addFavourite', options)
    return res.data
  }

  static async removeFavourite(options) {
    const res = await publicApi.post('/deleteFavourite', options)
    return res.data
  }

  static async getFavourite(id) {
    const res = await publicApi.post('/getFavourites', id)
    return res.data
  }

  static async getSearch(query = "", pageSize = 9) {
    console.log(query)
    const res = await publicApi.get(
      `/search?&${query}`
    )
    return res.data
  }

  static async getBook(id) {
    const res = await publicApi.post('/book', id)
    return res.data
  }

  static async addOrder(order) {
    const res = await publicApi.post('/createOrder', order)
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
