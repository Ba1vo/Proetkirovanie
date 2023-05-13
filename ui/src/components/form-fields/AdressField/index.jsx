import BaseField from '../BaseField'

import './style.scss'

const AdressField = ({ rules, ...rest }) => {
  const nameRegExp =
  /^[А-Яа-я0-9\s]{4,40}$/


  const patternObject = {
    value: nameRegExp,
    message: 'Введите корректный адрес',
  }

  if (rules) {
    rules.pattern = patternObject
  } else {
    rules = {
      pattern: patternObject,
    }
  }

  return (
    <BaseField type='text' rules={rules} {...rest}>
    </BaseField>
  )
}

export default AdressField
