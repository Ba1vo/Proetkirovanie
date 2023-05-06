import classNames from 'classnames'
import PropTypes from 'prop-types'

const CheckMark = ({ className }) => {
  return (
    <svg
      className={classNames('check-mark', className)}
      width='13'
      height='10'
      viewBox='0 0 13 10'
      fill='none'
      xmlns='http://www.w3.org/2000/svg'
    >
      <path
        d='M1 4.2L5.125 9L12 1'
        stroke='#BA0C0C'
        stroke-width='2'
        stroke-linecap='round'
        stroke-linejoin='round'
      />
    </svg>
  )
}

CheckMark.propTypes = {
  className: PropTypes.string,
}

export default CheckMark
