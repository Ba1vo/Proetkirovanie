import { useState, cloneElement } from 'react'

import PropTypes from 'prop-types'
import { Link } from 'react-router-dom'

import Container from '../../../components/Container'

import './style.scss'

const AuthTemplate = ({ title, linkLabel, linkTo, children }) => {
  const [serverError, setServerError] = useState(null)

  const errorElement = serverError ? (
    <div className='auth-page__error'>{serverError}</div>
  ) : null

  return (
    <div className='auth-page'>
      <Container className='auth-page__container'>
        <h2 className='auth-page__title'>{title}</h2>
        {cloneElement(children, { setServerError })}
        {errorElement}
        <Link className='auth-page__link' to={linkTo}>
          {linkLabel}
        </Link>
      </Container>
    </div>
  )
}

AuthTemplate.propTypes = {
  title: PropTypes.string,
  linkLabel: PropTypes.string,
  linkTo: PropTypes.string,
  children: PropTypes.node,
}

export default AuthTemplate
