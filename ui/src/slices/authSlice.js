import { createSlice, createAsyncThunk } from '@reduxjs/toolkit'

import AuthService from '../services/authService'

const initialState = {
  user: null,
  loading: false,
}

export const register = createAsyncThunk(
  'auth/register',
  async (user, { rejectWithValue }) => {
    try {
      const data = await AuthService.register(user)
      localStorage.setItem('token', data.access)
      return data.user
    } catch (error) {
      return rejectWithValue(error.response.data)
    }
  }
)

export const login = createAsyncThunk(
  'auth/login',
  async (user, { rejectWithValue }) => {
    try {
      const data = await AuthService.login(user)
      localStorage.setItem('token', data.access)
      return data.user
    } catch (error) {
      return rejectWithValue(error.response.data)
    }
  }
)

export const checkAuth = createAsyncThunk('auth/checkAuth', async () => {
  const data = await AuthService.checkAuth()
  localStorage.setItem('token', data.access)
  return data.user
})

export const logout = createAsyncThunk('auth/logout', async () => {
  await AuthService.logout()
  localStorage.removeItem('token')
})

const authSlice = createSlice({
  name: 'auth',
  initialState,
  extraReducers: (builder) => {
    builder
      .addCase(register.pending, (state) => {
        state.loading = true
      })
      .addCase(register.fulfilled, (state, action) => {
        state.user = action.payload
        state.loading = false
      })
      .addCase(register.rejected, (state) => {
        state.loading = false
      })
      .addCase(login.pending, (state) => {
        state.loading = true
      })
      .addCase(login.fulfilled, (state, action) => {
        state.user = action.payload
        state.loading = false
      })
      .addCase(login.rejected, (state) => {
        state.loading = false
      })
      .addCase(logout.fulfilled, (state) => {
        state.user = null
      })
      .addCase(checkAuth.fulfilled, (state, action) => {
        state.user = action.payload
      })
  },
})

export default authSlice.reducer
