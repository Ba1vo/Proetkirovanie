import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'
import { AiOutlineMenu, AiOutlineClose } from 'react-icons/ai'
import { Link } from 'react-router-dom'

import AccountIcon from '../icons/AccountIcon'

import './style.scss'

const Menu = ({ className, user, setUser}) => {
  const [menuOpen, setMenuOpen] = useState(false)

  const onLinkClick = (event) => {
    const tagName = event.target.tagName
    if (tagName === 'A' || tagName === 'SPAN') {
      setMenuOpen(false)
    }
  }

  const onMenuIconClick = () => {
    setMenuOpen((menuOpen) => !menuOpen)
  }

  return (
    <div className={classNames('menu', className)}>
      <nav className={menuOpen ? 'menu__body menu__body_open' : 'menu__body'}>
        <ul className='menu__list' onClick={onLinkClick}>
          <li className='menu__item'>
            <Link className='menu__link' to='/catalog'>
              Каталог
            </Link>
          </li>
          { user ?
          <li className='menu__item'>
            <Link className='menu__link' to='/favourites'>
              Избранное
            </Link>
          </li> : null
          }
          <li className='menu__item'>
            <Link className='menu__link' to='/cart'>
              Корзина
            </Link>
          </li>
          <li className='menu__item'>
            <Link className='menu__link' to='/login'>
              <AccountIcon className={user ?'menu__account-image__signed' :'menu__account-image__nonsigned'} />
              <span className='menu__account-label'> Аккаунт</span>
            </Link>
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
