import Container from '../Container'
import Logo from '../Logo'

import './style.scss'

const Footer = () => {
  return (
    <div className='footer'>
      <Container className='footer__container'>
        <div className='footer__logo-div'>
          <Logo className='footer__logo' color1='#808080' color2='#808080' />
        </div>
        <div className='footer__time'>
          <b className='footer__bold'>Режим работы:</b>
          <br />
          8:00 - 22:00 Без выходных.
        </div>
        <div className='footer__address'>
          <b className='footer__bold'>Адрес магазина:</b>
          <br />
          г. Томск ул. Пушкина 146
        </div>
        <div className='footer__contact'>
          <b className='footer__bold'>Контакты:</b>
          <br />
          +7 800 555 35 35
          <br />
          wine_fi@mail.ru
        </div>
      </Container>
    </div>
  )
}

export default Footer
