import { configureStore } from '@reduxjs/toolkit'

import authReducer from '../slices/authSlice'
import catalogProductsReducer from '../slices/catalogProductsSlice'
import discountedProductsReducer from '../slices/discountedProductsSlice'
import popularProductsReducer from '../slices/popularProductsSlice'

const store = configureStore({
  reducer: {
    auth: authReducer,
    discountedProducts: discountedProductsReducer,
    popularProducts: popularProductsReducer,
    catalogProducts: catalogProductsReducer,
  },
  devTools: process.env.NODE_ENV !== 'production',
})

export default store
