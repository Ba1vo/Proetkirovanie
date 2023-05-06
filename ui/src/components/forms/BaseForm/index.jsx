import { Children, cloneElement } from 'react'

import PropTypes from 'prop-types'
import { useForm } from 'react-hook-form'

import './style.scss'

const BaseForm = ({ onSubmit, children }) => {
  const {
    register,
    handleSubmit,
    reset,
    setError,
    formState: { errors },
  } = useForm()

  const submitHandler = (data) => {
    onSubmit(data, setError, reset)
  }

  return (
    <form className='form' onSubmit={handleSubmit(submitHandler)}>
      {Children.map(children, (child) => {
        const { name } = child.props
        if (name) {
          return cloneElement(child, { register, error: errors[name] })
        }
        return child
      })}
    </form>
  )
}

BaseForm.propTypes = {
  onSubmit: PropTypes.func,
  children: PropTypes.node,
}

export default BaseForm
