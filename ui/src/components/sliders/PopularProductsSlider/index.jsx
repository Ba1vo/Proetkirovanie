import { useEffect } from 'react'

import PropTypes from 'prop-types'
import { useDispatch, useSelector } from 'react-redux'

//import { getPopularProducts } from '../../../slices/popularProductsSlice'
import ProductSlider from '../ProductSlider/'
import ProductService from '../../../services/productService'
import { useState } from 'react'

const PopularProductsSlider = ({ className }) => {
  const [state, setState] = useState({products: null, inProgress: true}) 
  console.log(state)
  useEffect(() => {
    ProductService.getPopularBooks().then(
      (res) =>{
        setState({products: res, inProgress: false})
      }
    )
  }, [])

  return <ProductSlider className={className} loading={state.inProgress} products={state.products} topic='Популярное' />
}

PopularProductsSlider.propTypes = {
  className: PropTypes.string,
}

export default PopularProductsSlider
