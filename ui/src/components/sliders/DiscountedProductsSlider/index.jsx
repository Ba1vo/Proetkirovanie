import { useEffect, useState } from 'react'

import PropTypes from 'prop-types'
import { useDispatch, useSelector } from 'react-redux'

import { getDiscountedProducts } from '../../../slices/discountedProductsSlice'
import ProductSlider from '../ProductSlider/'
import ProductService from '../../../services/productService'

const DiscountedProductsSlider = ({ className }) => {
  const dispatch = useDispatch()
  const state = useSelector((state) => state.discountedProducts)

  useEffect(() => {
    dispatch(getDiscountedProducts())
  }, [dispatch])

  return <ProductSlider className={className} {...state} topic='Скидки' />
}

DiscountedProductsSlider.propTypes = {
  className: PropTypes.string,
}

export default DiscountedProductsSlider
