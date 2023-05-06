import { LoginForm } from '../../components/forms'
import AuthTemplate from '../templates/AuthTemplate'

const LoginPage = ({user, setUser}) => {
  return (
    <AuthTemplate
      title='Вход'
      linkLabel='Зарегистрироваться'
      linkTo='/registration'
    >
      <LoginForm user = {user} setUser = {setUser} />
    </AuthTemplate>
  )
}

export default LoginPage
