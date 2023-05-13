import { useEffect } from 'react'

import PropTypes from 'prop-types'
import { useDispatch, useSelector } from 'react-redux'

import { getPopularProducts } from '../../../slices/popularProductsSlice'
import ProductSlider from '../ProductSlider/'
import ProductService from '../../../services/productService'
import { useState } from 'react'

const PopularProductsSlider = ({ className }) => {
  const dispatch = useDispatch()
  const state = useSelector((state) => state.popularProducts)

  useEffect(() => {
    dispatch(getPopularProducts())
  }, [dispatch])

  return <ProductSlider className={className} {...state} topic='Популярное' />
}

PopularProductsSlider.propTypes = {
  className: PropTypes.string,
}

export default PopularProductsSlider
