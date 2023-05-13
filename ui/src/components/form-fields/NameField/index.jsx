import BaseField from '../BaseField'

import './style.scss'

const NameField = ({ rules, ...rest }) => {
  const nameRegExp =
  /^[А-Яа-я]{1,15}$/


  const patternObject = {
    value: nameRegExp,
    message: 'Разрешено до 15 русских букв',
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

export default NameField
