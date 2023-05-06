import { useState, useEffect } from 'react'

import classNames from 'classnames'
import { useSearchParams } from 'react-router-dom'

import Select from '../form-inputs/Select'

import './style.scss'

const options = [
  { label: 'По названию', value: 'name' },
  { label: 'По скидке', value: '-discount' },
  { label: 'Сначала дешёвые', value: 'price' },
  { label: 'Сначала дорогие', value: '-price' },
]

const ProductSort = ({ className }) => {
  const [searchParams, setSearchParams] = useSearchParams()
  const [option, setOption] = useState(options[0])

  const onOptionChange = (option) => {
    setOption(option)
    setSearchParams((searchParams) => {
      searchParams.set('order_by', option.value)
      searchParams.delete('page')
      return searchParams
    })
  }

  useEffect(() => {
    const option = options.find((option) => {
      return option.value === searchParams.get('order_by')
    })
    setOption(option ?? options[0])
  }, [searchParams])

  return (
    <div className={classNames('product-sort', className)}>
      <div className='product-sort__text'>Сортировать</div>
      <Select
        className='product-sort__select'
        options={options}
        onChange={onOptionChange}
        selectedOption={option}
      />
    </div>
  )
}

export default ProductSort
