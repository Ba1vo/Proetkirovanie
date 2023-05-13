import { useEffect, useState } from 'react'
import CartList from '../../components/CartList'
import Container from '../../components/Container'
import FavouritesList from '../../components/FavoritesProductList'
import OrdersList from '../../components/OrdersList'
import ProductFilter from '../../components/ProductFilter'
import ProductList from '../../components/ProductList'
import ProductSort from '../../components/ProductSort'
import Search from '../../components/Search'
import CartForm from '../../components/forms/CartForm'

import './style.scss'

const CartPage = () => {
  const [state, setState] = useState({books: [], ids: JSON.parse(localStorage.getItem('cart')), loading: true, error: null})
  useEffect(() => {
    if (state.ids === null){
      localStorage.setItem('cart', JSON.stringify([]))
      setState({...state, ids: []})
    }
  }, [])

  return (
    <div className='catalog-page'>
      <Container className='catalog-page__container'>
        <h2 className='catalog-page__title'> Корзина</h2>
        <div className='catalog-page__content'>
          <div className='catalog-page__products'>
            <CartList state={state} setState={setState} />
          </div>
          <CartForm state={state}></CartForm>
        </div>
      </Container>
    </div>
  )
}

export default CartPage
