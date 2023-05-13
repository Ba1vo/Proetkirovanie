import Container from '../../components/Container'
import FavouritesList from '../../components/FavoritesProductList'
import OrdersList from '../../components/OrdersList'
import ProductFilter from '../../components/ProductFilter'
import ProductList from '../../components/ProductList'
import ProductSort from '../../components/ProductSort'
import Search from '../../components/Search'

import './style.scss'

const OrderPage = () => {
  return (
    <div className='catalog-page'>
      <Container className='catalog-page__container'>
        <h2 className='catalog-page__title'> Заказы</h2>
        <div className='catalog-page__content'>
          <div className='catalog-page__products'>
            <OrdersList />
          </div>
        </div>
      </Container>
    </div>
  )
}

export default OrderPage
