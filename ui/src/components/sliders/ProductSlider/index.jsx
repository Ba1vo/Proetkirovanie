import classNames from 'classnames'
import PropTypes from 'prop-types'
import { Navigation } from 'swiper'

import ProductCard from '../../../components/ProductCard'
import ErrorIndicator from '../../ErrorIndicator'
import Spinner from '../../Spinner'
import Slider from '../Slider'

import './style.scss'

const ProductSlider = ({ products, topic, loading, error, className }) => {
  const config = {
    modules: [Navigation],
    navigation: true,
    slidesPerView: 1,
    slidesPerGroup: 1,
    speed: 800,
    breakpoints: {
      420: {
        slidesPerGroup: 1,
        slidesPerView: 1.5,
      },
      537: {
        slidesPerGroup: 2,
        slidesPerView: 2,
      },
      780: {
        slidesPerGroup: 3,
        slidesPerView: 3,
      },
      1035: {
        speed: 1000,
        slidesPerGroup: 4,
        slidesPerView: 4,
      },
    },
  }

  return (
    <div className={classNames('product-slider', className)}>
      <div className='product-slider__topic'>{topic}</div>
      <div
        className={classNames('product-slider__wrapper', {
          'product-slider__wrapper_not-loaded': loading || error,
        })}
      >
        {loading ? <Spinner /> : null}
        {error ? <ErrorIndicator /> : null}
        {!(loading || error) ? (
          <Slider
            className='product-slider__slider'
            items={products}
            config={config}
          >
            {(product) => <ProductCard {...product} />}
          </Slider>
        ) : null}
      </div>
    </div>
  )
}

ProductSlider.propTypes = {
  products: PropTypes.arrayOf(PropTypes.object),
  topic: PropTypes.string.isRequired,
  loading: PropTypes.bool,
  error: PropTypes.bool,
  className: PropTypes.string,
}

export default ProductSlider
