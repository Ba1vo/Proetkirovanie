import classNames from 'classnames'
import PropTypes from 'prop-types'

import './style.scss'

const BaseField = ({
  name,
  type,
  rules,
  register,
  error,
  required,
  children,
  className,
  ...rest
}) => {
  const requiredError = 'Обязательно для заполнения'

  if (rules && required) {
    rules.required = requiredError
  }

  const inputClasses = error
    ? 'field__input field__input_error'
    : 'field__input'

  const optionalElement = children ? (
    <div className='field__optional-element'>{children}</div>
  ) : null

  const errorElement = error ? (
    <div className='field__error'>{error?.message}</div>
  ) : null

  return (
    <div className={classNames('field', className)}>
      <div className='field__container'>
        <input
          className={inputClasses}
          type={type}
          {...rest}
          {...register(name, rules)}
        />
        {optionalElement}
      </div>
      {errorElement}
    </div>
  )
}

BaseField.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.oneOf(['text', 'email', 'password']).isRequired,
  rules: PropTypes.object,
  register: PropTypes.func.isRequired,
  error: PropTypes.object,
  required: PropTypes.bool,
  children: PropTypes.node,
  className: PropTypes.string,
}

export default BaseField
