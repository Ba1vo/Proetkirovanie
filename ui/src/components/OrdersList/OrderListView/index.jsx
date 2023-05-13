import Order from '../../Order'
import Pagination from '../../Pagination'
import ProductCard from '../../ProductCard'

import './style.scss'

const OrdersListView = ({ products, count }) => {
  console.log(count)
  if (count === 0) {
    return (
      <div className='order-list product-list_not-found'>
        Товаров не найдено
      </div>
    )
  }
  return (
    <div className='order-list'>
      <div className='order-list__products'>
        {products.map((product) => (
          <Order key={product.ID} {...product} />
        ))}
      </div>
    </div>
  )
}

export default OrdersListView
