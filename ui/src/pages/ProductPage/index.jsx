import { useParams } from 'react-router-dom'

import Container from '../../components/Container'
import Product from '../../components/Product'
import Search from '../../components/Search'

import './style.scss'

const ProductPage = () => {
  const { id } = useParams()
  return (
    <div className='product-page'>
      <Container className='product-page__container'>
        <Search />
        <Product
          className='product-page__info'
          id={id}
          image='../img/solntce_abhazii_red_04032022 1.svg'
          name='Солнце Абхазии'
          color='Красное'
          sugarType='Полусладкое'
          country='Россия'
          volume='0.75'
          actualPrice='1020'
          price='1120'
          discount='15'
          availability={true}
        />
      </Container>
    </div>
  )
}

export default ProductPage
