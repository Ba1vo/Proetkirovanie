import Pagination from '../../Pagination'
import ProductCard from '../../ProductCard'

import './style.scss'

const ProductList = ({ products, pageCount, forcePage, onPageClick, count }) => {
  console.log(count)
  if (count === 0) {
    return (
      <div className='product-list product-list_not-found'>
        Товаров не найдено
      </div>
    )
  }
  return (
    <div className='product-list'>
      <div className='product-list__products'>
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

export default ProductList
