import { useParams } from 'react-router-dom'

import Container from '../../components/Container'
import Product from '../../components/Product'
import Search from '../../components/Search'

import './style.scss'
import { useEffect, useState } from 'react'
import ProductService from '../../services/productService'

const ProductPage = () => {
  const { id } = useParams()
  const [state, setState] = useState({book: {}, inProgress: true})
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
          id={id}
          image={state.book.Photo}
          name={state.book.Name}
          price={state.book.Price}
          discount={state.book.Discont}
          isbn={state.book.ISBN}
          description={state.book.Desc}
          dimension ={state.book.Dimensions}
          authors ={state.book.Authors}
          publishers ={state.book.Publishers}
          genres ={state.book.Genres}
          amount ={state.book.Amount}
          date ={state.book.Date}
          availability={true}
        />
      </Container>
    </div>
  )
}

export default ProductPage
