import { useEffect, useState } from 'react'

import PropTypes from 'prop-types'
import { useDispatch, useSelector } from 'react-redux'

//import { getDiscountedProducts } from '../../../slices/discountedProductsSlice'
import ProductSlider from '../ProductSlider/'
import ProductService from '../../../services/productService'

const DiscountedProductsSlider = ({ className }) => {
  const [state, setState] = useState({products: null, inProgress: true})
  console.log(state)
  useEffect(() => {
    ProductService.getDiscountedBooks().then(
      (res) =>{
        setState({products: res, inProgress: false})
      }
    )
  }, [])

  return <ProductSlider className={className} loading={state.inProgress} products={state.products} topic='Скидки' />
}

DiscountedProductsSlider.propTypes = {
  className: PropTypes.string,
}

export default DiscountedProductsSlider
