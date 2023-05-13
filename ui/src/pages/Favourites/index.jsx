import Container from '../../components/Container'
import FavouritesList from '../../components/FavoritesProductList'
import ProductFilter from '../../components/ProductFilter'
import ProductList from '../../components/ProductList'
import ProductSort from '../../components/ProductSort'
import Search from '../../components/Search'

import './style.scss'

const FavouritesPage = () => {
  return (
    <div className='catalog-page'>
      <Container className='catalog-page__container'>
        <h2 className='catalog-page__title'> Избранное</h2>
        <div className='catalog-page__content'>
          <div className='catalog-page__products'>
            <FavouritesList />
          </div>
        </div>
      </Container>
    </div>
  )
}

export default FavouritesPage
