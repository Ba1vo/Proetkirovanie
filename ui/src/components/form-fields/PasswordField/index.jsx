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
  const passRegExp =
  /^(?=.*\d)(?=.*[A-Za-z])[A-Za-z\d]{5,16}$/


  const patternObject = {
    value: passRegExp,
    message: 'Разрешены латиница и цифр',
  }

  if (rules) {
    rules.pattern = patternObject
  } else {
    rules = {
      pattern: patternObject,
    }
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
