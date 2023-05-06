import TextField from '../TextField'

const NumberField = ({ name, rules, setValue, ...rest }) => {
  const onChange = (event) => {
    setValue(name, event.target.value.replace(/\D/g, ''))
  }

  if (rules) {
    rules.onChange = onChange
  } else {
    rules = { onChange }
  }

  return <TextField name={name} rules={rules} {...rest} />
}

export default NumberField
