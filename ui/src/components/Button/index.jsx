import classNames from 'classnames'
import PropTypes from 'prop-types'

import './style.scss'

const Button = ({ children, outlined, className, ...rest }) => {
  return (
    <button
      className={classNames('button', className, {
        'button_outlined': outlined,
      })}
      {...rest}
    >
      {children}
    </button>
  )
}

Button.propTypes = {
  type: PropTypes.string,
  disabled: PropTypes.bool,
  onClick: PropTypes.func,
  children: PropTypes.node,
  className: PropTypes.string,
}

Button.defaultProps = {
  type: 'button',
}

export default Button
