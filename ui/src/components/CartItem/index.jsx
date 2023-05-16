import { useEffect, useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'

import './style.scss'
import Counter from '../Counter'

const CartItem = ({
  ID,
  Photo,
  Name,
  Authors,
  Price,
  Discount,
  Count,
  Delete,
  state, 
  setState
}) => {

  const [count, setCount] = useState(Count)
  const [isLiked, setLike] = useState(false)

  useEffect(() => {
    let array = JSON.parse(localStorage.getItem('cart'))
    let bookIndex = array.findIndex((obj => obj.id == ID))
    if (bookIndex === -1) {
      setCount(1)
      return
    }

    setCount(array[bookIndex].amount)
  }, [])

  useEffect(() => {
    let array = JSON.parse(localStorage.getItem('cart'))
    let index = array.findIndex((obj => obj.id == ID))
    if (index !== -1){
      array[index].amount = count 
      setState({...state, ids: array})
      localStorage.setItem('cart', JSON.stringify(array));
    } 
  }, [count])

  let calcPrice = Math.round(Price*(100-Discount)/100*count).toFixed(2)
  return (
    <div className='cart-product__product'> 
    <div className='cart-product__img-div'>
    <img
      className='cart-product__img'
      src={Photo}
      alt='Изображение товара'
    />
    </div>
    <div className='cart-product__description'>
      {Name}.
      <div className='cart-product__author'>
        {Authors}.
      </div>
    </div>
    <div className='cart-product__counter'>
      <Counter count={count} setCount={setCount} />
    </div>
    <div className='cart-product__price-and-counter'>
      <div className='cart-product__price'>
        {calcPrice} ₽.
        {Discount ? (
          <span className='cart-product__old-price'>{Price*count} ₽</span>
        ) : null}
      </div>
    </div>
    </div>
  )
}

CartItem.propTypes = {
  imgSrc: PropTypes.string,
  type: PropTypes.string,
  name: PropTypes.string,
  color: PropTypes.string,
  sweetness: PropTypes.string,
  country: PropTypes.string,
  volume: PropTypes.number,
  price: PropTypes.number,
  Amount: PropTypes.number,
  added: PropTypes.bool,
}

export default CartItem
