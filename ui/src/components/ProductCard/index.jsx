import { useEffect, useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'

import Button from '../Button'
import Counter from '../Counter'
import LikeIcon from '../icons/LikeIcon'

import './style.scss'
import { set } from 'react-hook-form'
import ProductService from '../../services/productService'
import { useDispatch, useSelector } from 'react-redux'

const ProductCard = ({
  ID,
  Photo,
  Name,
  Authors,
  Price,
  Discount,
  Amount,
  Favourite
}) => {
  const [isAdded, setAdd] = useState(JSON.parse(localStorage.getItem('cart')).findIndex((obj => obj.id == ID)) !== -1)
  const [count, setCount] = useState(null)
  const [isLiked, setLike] = useState(Favourite)
  const user = useSelector((state) => state.auth.user)

  useEffect(() => {
    let array = JSON.parse(localStorage.getItem('cart'))
    let bookIndex = array.findIndex((obj => obj.id == ID))
    if (bookIndex === -1) {
      setAdd(false)
      return
    }
    setAdd(true)
    setCount(array[bookIndex].amount)
  }, [])

  const changeAdd = () => {
    let array = JSON.parse(localStorage.getItem('cart'))
    if (array === null){
       array = []
    }
    console.log(array)
    if (isAdded){
      array = array.filter(book => book.id != ID)
      localStorage.setItem('cart', JSON.stringify(array));
    } else{
      array.push({id: ID, amount: count})
      setCount(1)
      localStorage.setItem('cart', JSON.stringify(array));
    }
    setAdd((isAdded) => !isAdded)
  }

  useEffect(() => {
    let array = JSON.parse(localStorage.getItem('cart'))
    let index = array.findIndex((obj => obj.id == ID))
    if (index !== -1 && count !== null){
      array[index].amount = count 
      localStorage.setItem('cart', JSON.stringify(array));
    } 
  }, [count])

  const changeLike = () => {
    if (isLiked) {
      ProductService.deleteFavourite(ID).then(() =>{
        setLike(false)
       }).catch({
  
       })
    } else{
      ProductService.addFavourite(ID).then(() =>{
        setLike(true)
       }).catch({
  
       })
    }
  }

  let calcPrice = Discount ? Math.round(Price*(100-Discount)/100).toFixed(2) : Price
  return (
    <div className='product-card'>
      <a href={`/book/${ID}`}>
        <div className='product-card__a' />
      </a>

      <div className='product-card__img-div'>
        <img
          className='product-card__img'
          src={Photo}
          alt='Изображение товара'
        />
      </div>

      {Discount ? (
        <div className='product-card__discount'>-{Discount}%</div>
      ) : null}
      
      {Object.keys(user).length !== 0 ? <div onClick={changeLike}>
        <LikeIcon
          className={classNames('product-card__like', {
            'product-card__like_active': isLiked,
          })}
        />
      </div> : null
      }

      <div className='product-card__description'>
        {Name}.
      </div>
      <div className='product-card__author'>
        {Authors}.
      </div>
      <div className='product-card__price-and-counter'>
        <div className='product-card__price'>
          {calcPrice} ₽/шт.
          {Discount ? (
            <span className='product-card__old-price'>{Price} ₽/шт.</span>
          ) : null}
        </div>
        <div className='product-card__counter-div'>
          {isAdded ? <Counter count={count} setCount={setCount} /> : null}
        </div>
      </div>

      <Button
        className={classNames('product-card__btn', {
          'product-card__btn_in-basket': isAdded,
        })}
        onClick={changeAdd}
      >
        {isAdded ? 'В корзине' : 'В корзину'}
      </Button>
    </div>
  )
}

ProductCard.propTypes = {
  imgSrc: PropTypes.string,
  type: PropTypes.string,
  name: PropTypes.string,
  color: PropTypes.string,
  sweetness: PropTypes.string,
  country: PropTypes.string,
  volume: PropTypes.number,
  price: PropTypes.number,
  like: PropTypes.bool,
  added: PropTypes.bool,
}

export default ProductCard
