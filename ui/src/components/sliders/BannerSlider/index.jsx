import { useEffect } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'
import { useDispatch, useSelector } from 'react-redux'
import { Navigation, Pagination, Autoplay } from 'swiper'

//import { getBanners } from '../../../slices/bannersSlice'
import ErrorIndicator from '../../ErrorIndicator'
import Spinner from '../../Spinner'
import Slider from '../Slider'

import './style.scss'

const BannerSlider = ({ className }) => {
  const config = {
    modules: [Navigation, Pagination, Autoplay],
    loop: true,
    speed: 900,
    autoplay: {
      delay: 5000,
      disableOnInteraction: false,
    },
    slidesPerView: 1,
    navigation: true,
    pagination: { clickable: true },
  }

  const dispatch = useDispatch()
  const { banners, loading, error } = useSelector((state) => state.banners)

  useEffect(() => {
    //dispatch(getBanners())
  }, [dispatch])

  return (
    <div className={classNames('banner-slider', className)}>
      <div className='banner-slider__content'>
        {loading ? <Spinner /> : null}
        {error ? <ErrorIndicator /> : null}
        {!(loading || error) ? (
          <Slider
            className='banner-slider__slider'
            items={banners}
            config={config}
          >
            {(banner) => (
              <img
                className='banner-slider__img'
                src={banner.image}
                alt='Рекламный баннер'
              />
            )}
          </Slider>
        ) : null}
      </div>
    </div>
  )
}

BannerSlider.propTypes = {
  className: PropTypes.string,
}

export default BannerSlider
