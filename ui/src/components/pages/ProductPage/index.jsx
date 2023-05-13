import { useParams } from 'react-router-dom'

import Container from '../../components/Container'
import Product from '../../components/Product'
import Search from '../../components/Search'

import './style.scss'
import { useDispatch } from 'react-redux'
import { useState } from 'react'
import ProductService from '../../../services/productService'

const ProductPage = () => {
  const { id } = useParams()
  const [state, setState] = useState()
  console.log(22)
  useEffect(() => {
      ProductService.getBook(id).then(
        (res) =>{
          setState({book: res, inProgress: false})
        })
  }, [])

  return (
    <div className='product-page'>
      <Container className='product-page__container'>
        <Search />
        <Product
          className='product-page__info'
          availability={true}
        />
      </Container>
    </div>
  )
}

export default ProductPage
