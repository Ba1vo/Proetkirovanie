import Container from '../../components/Container'
import ProductFilter from '../../components/ProductFilter'
import ProductList from '../../components/ProductList'
import ProductSort from '../../components/ProductSort'
import Search from '../../components/Search'

import './style.scss'

const CatalogPage = () => {
  return (
    <div className='catalog-page'>
      <Container className='catalog-page__container'>
        <Search />
        <h2 className='catalog-page__title'>Каталог</h2>
        <div className='catalog-page__content'>
          <ProductFilter className='catalog-page__filter' />
          <div className='catalog-page__products'>
            <ProductSort className='catalog-page__sorter' />
            <ProductList className={'product-list-container'}/>
          </div>
        </div>
      </Container>
    </div>
  )
}

export default CatalogPage
