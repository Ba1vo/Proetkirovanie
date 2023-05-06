import classNames from 'classnames'
import PropTypes from 'prop-types'

import './style.scss'

const TextField = ({
  name,
  type = 'text',
  rules,
  register,
  error,
  required,
  small = false,
  children,
  className,
  ...rest
}) => {
  if (rules && required) {
    rules.required = 'Обязательно для заполнения'
  }

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
          className={classNames('field__input ', {
            'field__input_small': small,
            'field__input_error': error,
          })}
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

TextField.propTypes = {
  name: PropTypes.string.isRequired,
  type: PropTypes.oneOf(['text', 'email', 'password']),
  rules: PropTypes.object,
  register: PropTypes.func.isRequired,
  error: PropTypes.object,
  required: PropTypes.bool,
  small: PropTypes.bool,
  children: PropTypes.node,
  className: PropTypes.string,
}

export default TextField
