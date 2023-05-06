import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import ProductService from '../services/productService'

const initialState = {
  products: [],
  loading: true,
  error: false,
}

export const getPopularProducts = createAsyncThunk(
  'popularProducts/getProducts',
  async () => {
    return await ProductService.getPopularProducts()
  }
)

const popularProductsSlice = createSlice({
  name: 'popularProducts',
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(getPopularProducts.pending, (state) => {
        state.loading = true
        state.error = false
      })
      .addCase(getPopularProducts.fulfilled, (state, action) => {
        state.products = action.payload
        state.loading = false
      })
      .addCase(getPopularProducts.rejected, (state) => {
        state.loading = false
        state.error = true
      })
  },
})

export default popularProductsSlice.reducer
