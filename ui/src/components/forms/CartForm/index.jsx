import PropTypes from 'prop-types'
import { Link, useNavigate } from 'react-router-dom'

//import { login } from '../../../slices/authSlice'
import LoadingButton from '../../LoadingButton'
import BaseForm from '../BaseForm'
import AuthService from '../../../services/authService'
import { usePromiseTracker, trackPromise } from "react-promise-tracker"
import BaseField from '../../form-fields/BaseField'
import { login } from '../../../slices/authSlice'
import './style.scss'
import React, { useState } from 'react'
import LoadingIcon from '../../icons/LoadingIcon'
import { useDispatch, useSelector } from 'react-redux'
import AdressField from '../../form-fields/AdressField'
import ProductService from '../../../services/productService'
import { useEffect } from 'react'

const CartForm = ({ state }) => {
  const [loading, setLoading] = useState(false)
  const navigate = useNavigate()
  const user = useSelector((state) => state.auth.user)
  console.log(state)
  console.log(user)
  let sum = 0
  if (state.books.length > 0){
    for (let i = 0; i < state.ids.length; i++) {
      sum+= Number(Math.round(state.books[i].Price*(100-state.books[i].Discount)/100*state.ids[i].amount).toFixed(2))
    }
  }
  
  const onSubmit = async (Adress, setError) => {
    setLoading(true)
      ProductService.createOrder(Adress, state.ids).then(r => {
        console.log(r)
        setLoading(false)
        localStorage.removeItem('cart')
        navigate('/orders')
      }).catch((error) =>{
        console.log(error)
        setLoading(false)
        if (error.field) {
          setError(error.field, { type: 'custom', message: error.detail })
        } 
      })
  } 
  console.log(user)
  return (
    <div className={typeof user === "undefined" || state.ids.length === 0 ? 'order-form-container__disabled': 'order-form-container'}>
      <div className='order-form-container__sum'>
      Общая сумма: {sum} ₽
      </div>
      <BaseForm onSubmit={onSubmit}>
        <AdressField name='Adress' placeholder='Адрес доставки' required />
        <LoadingButton type='submit' loading={loading}>
          Заказать
        </LoadingButton>
      </BaseForm>
    </div>
  )
}

CartForm.propTypes = {
  setServerError: PropTypes.func,
}

export default CartForm
