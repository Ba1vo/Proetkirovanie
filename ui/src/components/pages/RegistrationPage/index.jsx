import { RegistrationForm } from '../../components/forms'
import AuthTemplate from '../templates/AuthTemplate'

const RegistrationPage = () => {
  return (
    <AuthTemplate
      title='Регистрация'
      linkLabel='Я уже зарегистрирован(a)'
      linkTo='/login'
    >
      <RegistrationForm />
    </AuthTemplate>
  )
}

export default RegistrationPage
