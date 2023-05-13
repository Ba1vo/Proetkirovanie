import CartItem from '../../CartItem'

import './style.scss'

const CartListView = ({ books, Delete, ids }) => {
  console.log(books)
  if (books.length === 0) {
    return (
      <div className='cart-list cart-list_not-found'>
        Товаров не найдено
      </div>
    )
  }
  return (
    <div className='cart-list'>
      <div className='cart-list__products'>
        {books.map((product) => (
          <CartItem key={product.ID} Delete={Delete} {...product} Count={ids.amount} />
        ))}
      </div>
    </div>
  )
}

export default CartListView
