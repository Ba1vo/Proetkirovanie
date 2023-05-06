import Container from '../../components/Container'
import Search from '../../components/Search'
import {
  BannerSlider,
  PopularProductsSlider,
  DiscountedProductsSlider,
} from '../../components/sliders'

import './style.scss'

const MainPage = ({user, setUser}) => {
  console.log(user)
  /*return(
    <div>

    </div>
  )*/
  return (
    <div className='main-page'>
      <Container className='main-page__container'>
        <Search className='main-page__search' />
        <DiscountedProductsSlider className='main-page__product-slider' />
        <PopularProductsSlider className='main-page__product-slider' />
      </Container>
    </div>
  ) //<BannerSlider className='main-page__banner-slider' />
  //<PopularProductsSlider className='main-page__product-slider' />
}

export default MainPage
