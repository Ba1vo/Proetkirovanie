import { useState } from 'react'

import eyeIcon from './eye-icon.svg'
import BaseField from '../BaseField'

import './style.scss'

const PasswordField = ({ rules, ...rest }) => {
  const [type, setType] = useState('password')
  const [eyeOpen, setEyeOpen] = useState(true)

  const onEyeClick = () => {
    if (type === 'password') {
      setType('text')
      setEyeOpen(false)
    } else {
      setType('password')
      setEyeOpen(true)
    }
  }

  const passwordLengthError = 'Длина пароля от 4 до 12 символов'

  const rulesObject = {
    minLength: {
      value: 4,
      message: passwordLengthError,
    },
    maxLength: {
      value: 12,
      message: passwordLengthError,
    },
  }

  if (rules) {
    rules.minLength = rulesObject.minLength
    rules.maxLength = rulesObject.maxLength
  } else {
    rules = rulesObject
  }

  return (
    <BaseField className='password-field' type={type} rules={rules} {...rest}>
      <svg className='password-field__eye-icon' onClick={onEyeClick}>
        <use xlinkHref={`${eyeIcon}#${eyeOpen}`} />
      </svg>
    </BaseField>
  )
}

export default PasswordField
