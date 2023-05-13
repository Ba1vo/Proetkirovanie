import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import ProductService from '../services/productService'

const initialState = {
  products: [],
  loading: true,
  error: false,
}

export const getOrdersProducts = createAsyncThunk(
  'ordersProducts/getProducts',
  async () => {
    return await ProductService.getOrders()
  }
)

const ordersProductsSlice = createSlice({
  name: 'ordersProducts',
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(getOrdersProducts.pending, (state) => {
        state.loading = true
        state.error = false
      })
      .addCase(getOrdersProducts.fulfilled, (state, action) => {
        state.products = action.payload
        state.loading = false
      })
      .addCase(getOrdersProducts.rejected, (state) => {
        state.loading = false
        state.error = true
      })
  },
})

export default ordersProductsSlice.reducer
