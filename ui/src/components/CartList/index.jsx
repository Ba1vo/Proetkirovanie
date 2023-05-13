import { useEffect, useState } from 'react'

import classNames from 'classnames'
import { useDispatch, useSelector } from 'react-redux'
import { useMediaQuery } from 'react-responsive'
import { useNavigate, useSearchParams } from 'react-router-dom'

import { getOrdersProducts } from '../../slices/orderProductSlice'
import ErrorIndicator from '../ErrorIndicator'
import Spinner from '../Spinner'
import { useFocusEffect } from '@react-navigation/native';

import './style.scss'
import CartListView from './CartListView'
import ProductService from '../../services/productService'

const CartList = ({state, setState}) => {
  const isMobile = useMediaQuery({ query: '(max-width: 1080px)' })
  //localStorage.setItem('testObject', JSON.stringify(testObject));
  console.log(state.ids)
  useEffect(() => {
    if (state.ids.length === 0){
      setState({...state, loading: false})
      return
    }
    let ids = []
    state.ids.forEach(obj => ids.push(obj.id))
    ProductService.getCart(ids).then(r => {
      console.log(r)
      setState({...state, books: r, loading: false})
      console.log(state)
    }).catch((error) =>{
      console.log(error)
      setState({...state, loading: false, error: error})
      console.log(state)
    })
  }, [])

  const Delete = (id) => {
    let array = JSON.parse(localStorage.getItem('cart'))
    if (array === null){
       return
    }
    console.log(array)
    array = array.filter(book => book.id != id)
    let books = state.books.filter(book => book.ID != id)
    setState({...state, books: books, ids: array})
    localStorage.setItem('cart', JSON.stringify(array));
  }

  return (
    <div
      className={classNames('cart-list-container', {
        'cart-list-container_not-loaded': state.loading || state.error,
      })}
    >
      {state.loading ? <Spinner /> : null}
      {state.error ? <ErrorIndicator te /> : null}
      {!(state.loading || state.error) ? (
        <CartListView
          books={state.books}
          ids = {state.ids}
          Delete={Delete}
        />
      ) : null}
    </div>
  )
}

export default CartList
