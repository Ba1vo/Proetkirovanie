import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'
import { AiOutlineMenu, AiOutlineClose } from 'react-icons/ai'
import { Link } from 'react-router-dom'

import AccountIcon from '../icons/AccountIcon'

import './style.scss'
import AuthService from '../../services/authService'
import { useDispatch, useSelector } from 'react-redux'
import { logout } from '../../slices/authSlice'

const Menu = ({ className }) => {
  const [menuOpen, setMenuOpen] = useState(false)
  const [accountOpen, setAccountOpen] = useState(false)
  const dispatch = useDispatch()  
  const user = useSelector((state) => state.auth.user)
  console.log(user)
  const onLinkClick = (event) => {
    const tagName = event.target.tagName
    if (tagName === 'A' || tagName === 'SPAN') {
      setMenuOpen(false)
    }
  }

  const onMenuIconClick = () => {
    setMenuOpen((menuOpen) => !menuOpen)
  }

  const onAccountClick = () => {
    setAccountOpen((accountOpen) => !accountOpen)
  }

  const Exit = async () => {
    dispatch(logout()).unwrap()
  }
  console.log(Object.keys(user).length === 0)
  const AccountContent = 
  Object.keys(user).length === 0 ? 
      <Link className='menu__link' to='/login'>
      <AccountIcon className={'menu__account-image__nonsigned'} />
      <span className='menu__account-label'> Аккаунт</span>
      </Link> 
  : accountOpen ? 
      <Link className='menu__link' onClick={Exit} onMouseLeave={onAccountClick}>
        Выйти?
      </Link> :             
      <Link className='menu__link' onClick={onAccountClick}>
        <AccountIcon className={'menu__account-image__signed' } />
        {user.Name}
        <span className='menu__account-label'> Аккаунт</span>
      </Link>

  return (
    <div className={classNames('menu', className)}>
      <nav className={menuOpen ? 'menu__body menu__body_open' : 'menu__body'}>
        <ul className='menu__list' onClick={onLinkClick}>
          <li className='menu__item'>
            <Link className='menu__link' to='/catalog'>
              Каталог
            </Link>
          </li>
          { Object.keys(user).length > 0 ?
          <li className='menu__item'>
            <Link className='menu__link' to='/favourites'>
              Избранное
            </Link>
          </li> : null
          }
          { Object.keys(user).length > 0 ?
          <li className='menu__item'>
            <Link className='menu__link' to='/orders'>
              Заказы
            </Link>
          </li> : null
          }
          <li className='menu__item'>
            <Link className='menu__link' to='/cart'>
              Корзина
            </Link>
          </li>
          <li className='menu__item__acc'>
            {AccountContent}
          </li>
        </ul>
      </nav>
      <div className='menu__icon-container' onClick={onMenuIconClick}>
        {menuOpen ? (
          <AiOutlineClose className='menu__icon-close' />
        ) : (
          <AiOutlineMenu className='menu__icon-open' />
        )}
      </div>
    </div>
  )
}

Menu.propTypes = {
  className: PropTypes.string,
}

export default Menu
