import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'

import Button from '../Button'
import Counter from '../Counter'
import LikeIcon from '../icons/LikeIcon'

import './style.scss'

const OrderProduct = ({
  ID,
  Photo,
  Name,
  Authors,
  Price,
  Discount,
  Amount
}) => {
  const [isAdded, setAdd] = useState(false)
  const [isLiked, setLike] = useState(false)

  const changeAdd = () => {
    setAdd((isAdded) => !isAdded)
  }

  const changeLike = () => {
    setLike((isLiked) => !isLiked)
  }

  let calcPrice = Math.round(Price*(100-Discount)/100*Amount).toFixed(2)
  return (
    <div className='order-product__product'> 
    <div className='order-product__img-div'>
    <img
      className='order-product__img'
      src={Photo}
      alt='Изображение товара'
    />
    </div>
    <div className='order-product__description'>
      {Name}.
      <div className='order-product__author'>
        {Authors}.
      </div>
    </div>
    <div className='order-product__amount'>{Amount} шт.</div>
    <div className='order-product__price-and-counter'>
      <div className='order-product__price'>
        {calcPrice} ₽.
        {Discount ? (
          <span className='order-product__old-price'>{Price*Amount} ₽</span>
        ) : null}
      </div>
    </div>
    </div>
  )
}

OrderProduct.propTypes = {
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

export default OrderProduct
