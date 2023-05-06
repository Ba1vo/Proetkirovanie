import classNames from 'classnames'
import PropTypes from 'prop-types'

import LoupeIcon from '../icons/LoupeIcon'
import { getProducts } from '../../slices/catalogProductsSlice'

import './style.scss'
import { useNavigate, useSearchParams } from 'react-router-dom'
import { useEffect, useState } from 'react'
import { useDispatch } from 'react-redux'

const Search = ({ className }) => {
  const [searchParams, setSearchParams] = useSearchParams();  
  const [state, setState] = useState("")
  const dispatch = useDispatch()
  const navigate = useNavigate()
  const SearchChange = (e) => {
    setState(e.target.value)
  }

  useEffect(() => {
    setState(searchParams.get("str"))
  }, [])

  const Search = async (e) => {
    e.preventDefault()
    navigate(`/catalog?&page_size=9&str=${state ? state : ""}`)
  }

  return (
    <div className={classNames('search', className)}>
      <input
        className='search__input'
        type='text'
        placeholder='Поиск в каталоге'
        value={state || ""}
        onChange={SearchChange}
      />
      <button className='search__button' onClick={Search}>
        <LoupeIcon className='search__loupe' />
      </button>
    </div>
  )
}

Search.propTypes = {
  className: PropTypes.string,
}

export default Search
