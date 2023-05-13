import PropTypes from 'prop-types'
import { useNavigate } from 'react-router-dom'

//import { login } from '../../../slices/authSlice'
import { EmailField, PasswordField } from '../../form-fields'
import LoadingButton from '../../LoadingButton'
import BaseForm from '../BaseForm'
import AuthService from '../../../services/authService'
import { usePromiseTracker, trackPromise } from "react-promise-tracker"
import BaseField from '../../form-fields/BaseField'
import eyeIcon from './eye-icon.svg'
import { login } from '../../../slices/authSlice'

import React, { useState } from 'react'
import LoadingIcon from '../../icons/LoadingIcon'
import { useDispatch, useSelector } from 'react-redux'

const LoginForm = ({ user, setUser, setServerError }) => {
  /*
  const { promiseInProgress } = usePromiseTracker();
  const [type, setType] = useState("password");
  const [eyeOpen, setEyeOpen ] = useState(true);
  const [emailError, setEmailError ] = useState(null);
  const [passError, setPassError ] = useState(null);
  const [requestError, setRequestError ] = useState(null);
  const navigate = useNavigate()
  console.log(eyeOpen)
  console.log(type)
  const onEyeClick = () => {
    if (type === 'password') {
      setType('text')
      setEyeOpen(false)
    } else {
      setType('password')
      setEyeOpen(true)
    }
  }

  const onSubmit = async (e) => {
    e.preventDefault()
    let user = {
      Email: e.target.email.value,
      Pass: e.target.password.value
    }
    console.log(user)
    try {
      let response = await trackPromise(AuthService.login(user))
      setUser({
        ID: response.ID,
        Name:  response.Name,
        Email: response.Email,
        Date:  response.Date,
        Role:  response.Role
      })
      navigate('/')
    } catch (error) {
      console.log(error)
      if (error.response.data) {
        setRequestError({message: error.response.data})
      } else { 
        setRequestError({message:'Непредвиденная ошибка'})//
      }
    }
  }
  const emailRegExp =
  /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
  
  const passRegExp =
  /^(?=.*\d)(?=.*[A-Za-z])[A-Za-z\d]{5,16}$/

  const checkRegex = (e) => {
    console.log(e.target.value.match(emailRegExp))
    if (e.target.value.match(emailRegExp) === null){
      setEmailError({message: "Введите правильную почту"})
    } else {
      setEmailError(null)
    }
  }

  const passCheck = (e) => {
    console.log(e.target.value.match(emailRegExp))
    if (e.target.value.match(passRegExp) === null){
      setPassError({message: "Допустима латиница и цифры"})
    }else{
      setPassError(null)
    }
    if ((e.target.value.length < 5 || e.target.value.length > 16) && e.target.value.match(passRegExp) === null ){
      setPassError({message: "Пароль должен быть от 5 до 16 символов Допустима латиница и цифры"}) // XUETA
    }
  }

  const emailErrorElement = emailError ? (
    <div className='field__error'>{emailError?.message}</div>
  ) : null

  const passErrorElement = passError ? (
    <div className='field__error'>{passError?.message}</div>
  ) : null
  
  return (
    <React.Fragment> 
    <form className='form' onSubmit={onSubmit}>
    <div className='field'>
      <div className='field__container'>
        <input
        type = 'text'
        className='field__input'
        name = 'email'
        placeholder='Электронная почта'
        onBlur={checkRegex}
        />
      </div>
      {emailErrorElement}
    </div>
    <div className='field password-field'>
      <div className='field__container'>
        <input
        type = {type}
        className='field__input'
        name = 'password'
        placeholder='Пароль'
        onBlur={passCheck}
        />
      <div className='field__optional-element'>
        <svg className='password-field__eye-icon' onClick={onEyeClick}>
          <use xlinkHref={`${eyeIcon}#${eyeOpen}`} />
        </svg>
      </div>
      </div>
      {passErrorElement}
    </div>
    <button className="button loading-button" type="submit" disabled={promiseInProgress}>  {promiseInProgress ? <LoadingIcon className='loading-button__spinner' /> : "Войти"}</button>
    </form>
    {requestError ? <div class="auth-page__error"> {requestError.message} </div> : null}
    </React.Fragment>
/*
    <BaseForm onSubmit={onSubmit}>
      <EmailField name='email' placeholder='Электронная почта' required />
      <PasswordField name='password' placeholder='Пароль' required />
      <LoadingButton type='submit' loading={loading}>
        Войти
      </LoadingButton>
    </BaseForm>
    
   
  )*/
  const dispatch = useDispatch()
  const loading = useSelector((state) => state.auth.loading)
  const navigate = useNavigate()
    
  const onSubmit = async (user, setError) => {
    try {
      await dispatch(login(user)).unwrap()
      navigate('/')
    } catch (error) {
      if (error.field) {
        setError(error.field, { type: 'custom', message: error.detail })
      } else {
        setServerError('Непредвиденная ошибка')
      }
    }
  }

  return (
    <BaseForm onSubmit={onSubmit}>
      <EmailField name='Email' placeholder='Электронная почта' required />
      <PasswordField name='Pass' placeholder='Пароль' required />
      <LoadingButton type='submit' loading={loading}>
        Войти
      </LoadingButton>
    </BaseForm>
  )
}

LoginForm.propTypes = {
  setServerError: PropTypes.func,
}

export default LoginForm
