import { useState } from 'react'

import eyeIcon from './eye-icon.svg'
import TextField from '../TextField'

import './style.scss'

const PasswordField = ({ rules, ...rest }) => {
  const [type, setType] = useState('password')
  const [eyeOpen, setEyeOpen] = useState(true)

  const onEyeClick = () => {
    setEyeOpen((eyeOpen) => !eyeOpen)
    if (type === 'password') {
      setType('text')
    } else {
      setType('password')
    }
  }

  const passwordLengthError = 'Длина пароля от 4 до 12 символов'

  const passwordRules = {
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
    rules = { ...rules, ...passwordRules }
  } else {
    rules = passwordRules
  }

  return (
    <TextField className='password-field' type={type} rules={rules} {...rest}>
      <svg className='password-field__eye-icon' onClick={onEyeClick}>
        <use xlinkHref={`${eyeIcon}#${eyeOpen}`} />
      </svg>
    </TextField>
  )
}

export default PasswordField
