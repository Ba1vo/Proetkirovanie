import Pagination from '../../Pagination'
import ProductCard from '../../ProductCard'

import './style.scss'

const FavouritesListView = ({ products, pageCount, forcePage, onPageClick, count }) => {
  if (count === 0) {
    return (
      <div className='favourites-list favourites-list_not-found'>
        Нет товаров в избранном
      </div>
    )
  }
  console.log(products)
  return (
    <div className='favourites-list'>
      <div className='favourites-list__products'>
        {products.map((product) => (
          <ProductCard key={product.ID} {...product} />
        ))}
      </div>
      <Pagination
        forcePage={forcePage}
        pageCount={pageCount}
        onPageChange={onPageClick}
      />
    </div>
  )
}

export default FavouritesListView
