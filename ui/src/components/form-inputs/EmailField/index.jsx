import TextField from '../TextField'

const EmailField = ({ rules, ...rest }) => {
  const emailRegExp =
    /^((?=[a-zA-Z0-9])[a-zA-Z0-9!#$%&\\'*+\-/=?^_`.{|}~]{1,25})@(([a-zA-Z0-9-]){1,25}\.)([a-zA-Z0-9]{2,4})$/

  const patternRule = {
    value: emailRegExp,
    message: 'Введите корректную почту',
  }

  if (rules) {
    rules.pattern = patternRule
  } else {
    rules = { pattern: patternRule }
  }

  return <TextField rules={rules} {...rest} />
}

export default EmailField
