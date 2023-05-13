import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'

import Button from '../Button'
import Counter from '../Counter'
import LikeIcon from '../icons/LikeIcon'

import './style.scss'
import OrderProduct from '../OrderProduct'

const Order = ({
  ID,
  Date,
  Adress,
  Status,
  Books,
}) => {
  let sum = Number(0)
  return (
      <div className='order'>
          <div className='order__name'>
          Заказ №{ID} от {Date.slice(0,10)}. Статус: {Status}
          </div>
          {Books.map((book) =>(
          <OrderProduct key={book.ID} {...book} />
        ))}
        <div className='order__sum'>
          Сумма заказа: {Books.reduce((s, b) => s=s+Number(Math.round(Number(b.Price)*(100-Number(b.Discount))/100*Number(b.Amount)).toFixed(2)), 0)} ₽
          </div>
      </div>
  )
}

Order.propTypes = {
  imgSrc: PropTypes.string,
  type: PropTypes.string,
  name: PropTypes.string,
  color: PropTypes.string,
  sweetness: PropTypes.string,
  country: PropTypes.string,
  sum: PropTypes.number,
  price: PropTypes.number,
  like: PropTypes.bool,
  added: PropTypes.bool,
}

export default Order
