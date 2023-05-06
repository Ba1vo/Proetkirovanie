import classNames from 'classnames'
import PropTypes from 'prop-types'

import './style.scss'

const Container = ({ children, className }) => {
  return <div className={classNames('container', className)}>{children}</div>
}

Container.propTypes = {
  children: PropTypes.node,
  className: PropTypes.string,
}

export default Container
