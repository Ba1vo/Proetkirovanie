import { useState } from 'react'

import classNames from 'classnames'

import Button from '../Button'
import Counter from '../Counter'
import CheckMark from '../icons/CheckMark'
import LikeIcon from '../icons/LikeIcon'

import './style.scss'

const Product = ({
  id,
  image,
  name,
  price,
  discount,
  isbn,
  description,
  dimension,
  authors,
  publishers,
  genres,
  amount,
  date,
  className,
  availability,
}) => {
  const [isAdded, setAdd] = useState(false)
  const [isLiked, setLike] = useState(false)

  const changeAdd = () => {
    setAdd((isAdded) => !isAdded)
  }
  const actualPrice = discount ? Math.round(price*(100-discount)/100).toFixed(2) : price

  const changeLike = () => {
    setLike((isLiked) => !isLiked)
  }
  console.log(image)
  return (
    <div className={classNames('product', className)}>
      <div className='product__article'>Артикул {id}</div>
      <div className='product__content'>
        <div className='product__picture'>
          <img className='product__img' src={"/"+image} alt='Изображение товара' />

          {discount ? (
            <div className='product__discount'>-{discount}%</div>
          ) : null}

          <div onClick={changeLike}>
            <LikeIcon
              className={classNames('product__like', {
                'product__like_active': isLiked,
              })}
            />
          </div>
        </div>
        <div className='product__info'>
          <div className='product__name'>{name}</div>
          {authors}
          <div className='product__discription'>
            <span className='product__feature'>ISBN:</span>
            {isbn} <br/>
            <span className='product__feature'>Размеры:</span>
            {dimension}<br/>
            <span className='product__feature'>Издательства:</span>
            {publishers}<br/>
            <span className='product__feature'>Жанры:</span>
            {genres}<br/>
            <span className='product__feature'>Дата публикации:</span>
            {date ? date.slice(0,10) : "-"}<br/>
          </div>
        </div>

        <div className='product__buy'>
          <span className='product__feature'>Цена:</span>

          <div className='product__price'>
            {actualPrice} ₽/шт.
            {discount ? (
              <span className='product__old-price'>{price} ₽/шт.</span>
            ) : null}
          </div>

          <div className='product__avail-and-count'>
            <div className='product__availability'>
              {availability ? (
                <div>
                  <CheckMark /> В наличии
                </div>
              ) : (
                'Нет в наличии'
              )}
            </div>

            <div className='product__counter-div'>
              {isAdded ? <Counter /> : null}
            </div>
          </div>

          <Button
            className={classNames('product__btn', {
              'product__btn_in-basket': isAdded,
            })}
            onClick={changeAdd}
          >
            {isAdded ? 'В корзине' : 'В корзину'}
          </Button>
        </div>
      </div>
    </div>
  )
}

export default Product
