import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import ProductService from '../services/productService'

const initialState = {
  products: {
    products: [],
    count: 0,
    loading: true,
    error: false,
  },
  filters: {
    filters: {},
    loading: true,
    error: false,
  },
}

export const getProducts = createAsyncThunk(
  'catalogProducts/getSearch',
  async ({ queryString, pageSize }) => {
    return await ProductService.getSearch(queryString, pageSize)
  }
)

export const getFilters = createAsyncThunk(
  'catalogProducts/getFilters',
  async () => {
    return await ProductService.getFilters()
  }
)

const discountedProductsSlice = createSlice({
  name: 'catalogProducts',
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(getProducts.pending, (state) => {
        state.products.loading = true
        state.products.error = false
      })
      .addCase(getProducts.fulfilled, (state, action) => {
        console.log(action.payload)
        state.products.products = action.payload
        state.products.count = action.payload.length
        state.products.loading = false
      })
      .addCase(getProducts.rejected, (state) => {
        state.products.loading = false
        state.products.error = true
      })
      .addCase(getFilters.pending, (state) => {
        state.filters.loading = true
        state.filters.error = false
      })
      .addCase(getFilters.fulfilled, (state, action) => {
        state.filters.filters = action.payload
        state.filters.loading = false
      })
      .addCase(getFilters.rejected, (state) => {
        state.filters.loading = false
        state.filters.error = true
      })
  },
})

export default discountedProductsSlice.reducer
