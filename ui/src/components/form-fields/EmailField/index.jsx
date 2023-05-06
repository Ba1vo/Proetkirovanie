import BaseField from '../BaseField'

const EmailField = ({ rules, ...rest }) => {
  const emailRegExp =
    /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/

  const patternObject = {
    value: emailRegExp,
    message: 'Введите корректную почту',
  }

  if (rules) {
    rules.pattern = patternObject
  } else {
    rules = {
      pattern: patternObject,
    }
  }

  return <BaseField type='text' rules={rules} {...rest} />
}

export default EmailField
