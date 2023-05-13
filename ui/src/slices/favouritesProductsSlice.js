import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import ProductService from '../services/productService'

const initialState = {
    products: [],
    count: 0,
    loading: true,
    error: false,
}

export const getFavouriteProducts = createAsyncThunk(
  'favouriteProducts/getFavourites',
  async ({queryString, pageSize }) => {
    return await ProductService.getFavourite(pageSize, queryString)
  }
)

const favouritesProductsSlice = createSlice({
  name: 'favouriteProducts',
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(getFavouriteProducts.pending, (state) => {
        state.loading = true
        state.error = false
      })
      .addCase(getFavouriteProducts.fulfilled, (state, action) => {
        state.loading = false
        state.products = action.payload
        state.count = String(action.payload) === "null" ? 0 : action.payload.length
      })
      .addCase(getFavouriteProducts.rejected, (state) => {
        state.loading = false
        state.error = true
      })
  },
})

export default favouritesProductsSlice.reducer
