import classNames from 'classnames'
import PropTypes from 'prop-types'

import './style.scss'

const Checkbox = ({ name, label, register, rules, className, ...rest }) => {
  return (
    <label className={classNames('checkbox', className)}>
      <input type='checkbox' {...rest} {...register(name, rules)} />
      <span>{label}</span>
    </label>
  )
}

Checkbox.propTypes = {
  name: PropTypes.string.isRequired,
  label: PropTypes.oneOfType([PropTypes.string, PropTypes.number]),
  register: PropTypes.func.isRequired,
  rules: PropTypes.object,
  className: PropTypes.string,
}

export default Checkbox
