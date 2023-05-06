import Container from '../Container'
import Logo from '../Logo'
import Menu from '../Menu'

import './style.scss'

const Header = ({user, setUser}) => {
  return (
    <header className='header'>
      <Container className='header__container'>
        <Logo />
        <Menu className='header__menu' user={user} setUser={setUser} />
      </Container>
    </header>
  )
}

export default Header
