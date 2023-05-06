import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'
import { AiOutlineClose } from 'react-icons/ai'

import Accordion from '../../Accordion'
import Button from '../../Button'
import CheckboxList from '../../CheckboxList'
import CheckboxSearchList from '../../CheckboxSearchList'
import FilterRange from '../../FilterRange'

import './style.scss'

const ProductFilter = ({
  filters,
  register,
  setValue,
  watch,
  onSubmit,
  onReset,
}) => {
  const { minPrice, maxPrice, color, sugarType, country, volume } = filters
  const [filterOpen, setFilterOpen] = useState(false)

  const onFilterClick = () => {
    setFilterOpen((filterOpen) => !filterOpen)
  }

  return (
    <form className='product-filter' onSubmit={onSubmit}>
      <div className='product-filter__mobile-text' onClick={onFilterClick}>
        Фильтры
      </div>
      <div
        className={classNames('product-filter__wrapper', {
          'product-filter__wrapper_open': filterOpen,
        })}
      >
        <div className='product-filter__icon-container' onClick={onFilterClick}>
          {filterOpen ? (
            <AiOutlineClose className='product-filter__icon-close' />
          ) : null}
        </div>
        <Accordion title='Цена, ₽' defaultOpen>
          <FilterRange
            min={minPrice}
            max={maxPrice}
            minInputName='min_price'
            maxInputName='max_price'
            setValue={setValue}
            register={register}
          />
        </Accordion>
        <Accordion title='Страна' defaultOpen>
          <CheckboxSearchList
            name='country'
            searchInputName='country_term'
            watch={watch}
            values={country}
            register={register}
          />
        </Accordion>
        <Accordion title='Цвет' defaultOpen>
          <CheckboxList name='color' values={color} register={register} />
        </Accordion>
        <Accordion title='Сладость' defaultOpen>
          <CheckboxList
            name='sugar_type'
            values={sugarType}
            register={register}
          />
        </Accordion>
        <Accordion title='Объём, л'>
          <CheckboxList name='volume' values={volume} register={register} />
        </Accordion>
        <Button className='product-filter__btn' type='submit'>
          Применить
        </Button>
        <Button className='product-filter__btn' onClick={onReset} outlined>
          Сбросить
        </Button>
      </div>
    </form>
  )
}

ProductFilter.propTypes = {
  className: PropTypes.string,
}

export default ProductFilter
