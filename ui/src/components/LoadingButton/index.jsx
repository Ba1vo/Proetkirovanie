import classNames from 'classnames'
import PropTypes from 'prop-types'

import Button from '../Button'
import LoadingIcon from '../icons/LoadingIcon'

import './style.scss'

const LoadingButton = ({ loading, children, className, ...rest }) => {
  return (
    <Button
      className={classNames('loading-button', className)}
      disabled={loading}
      {...rest}
    >
      {loading ? <LoadingIcon className='loading-button__spinner' /> : children}
    </Button>
  )
}

LoadingButton.propTypes = {
  loading: PropTypes.bool,
  children: PropTypes.node,
  className: PropTypes.string,
}

export default LoadingButton
