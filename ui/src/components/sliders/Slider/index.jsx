import PropTypes from 'prop-types'
import { Swiper, SwiperSlide } from 'swiper/react'

import 'swiper/scss'
import 'swiper/scss/navigation'
import 'swiper/scss/pagination'

import './style.scss'

const Slider = ({ items, config = {}, children, className }) => {
  return (
    <Swiper className={className} {...config}>
      {items.map((item) => (
        <SwiperSlide key={item.ID}>{children(item)}</SwiperSlide>
      ))}
    </Swiper>
  )
}

Slider.propTypes = {
  items: PropTypes.arrayOf(PropTypes.object),
  config: PropTypes.object,
  children: PropTypes.func,
  className: PropTypes.string,
}

export default Slider
