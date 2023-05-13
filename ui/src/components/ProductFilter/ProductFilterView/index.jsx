import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'
import { AiOutlineClose } from 'react-icons/ai'
import Slider from '@mui/material/Slider';
import Box from '@mui/material/Box';

import Accordion from '../../Accordion'
import Button from '../../Button'
import CheckboxList from '../../CheckboxList'
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
  const { MinPrice, MaxPrice, Genres, MinDate, MaxDate } = filters
  const [filterOpen, setFilterOpen] = useState(false)
  const [dates, setDates] = useState([Number(MinDate.slice(0,4)), Number(MaxDate.slice(0,4))])
  const onFilterClick = () => {
    setFilterOpen((filterOpen) => !filterOpen)
  } 

  function valuetext(value) {
    return `${value}г.`;
  }
  const onDateChange = (event, value) => {
    setDates(value)
    setValue('dates', value)
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
            min={MinPrice}
            max={MaxPrice}
            minInputName='min_price'
            maxInputName='max_price'
            setValue={setValue}
            register={register}
          />
        </Accordion>
        <Accordion title='Дата печати'>
        <div className='accordion__title'>{dates[0]} 
          <span style={{float:"right"}}>
            {dates[1]}
          </span>
        </div>
        <Box px={2.5} py={1}>
        <Slider 
          disableSwap = {true}
          getAriaLabel={() => 'Date'}
          name='dates' 
          value={dates}
          onChange={onDateChange}
          min={Number(MinDate.slice(0,4))} 
          max={Number(MaxDate.slice(0,4))}
          valueLabelDisplay="off"
          getAriaValueText={valuetext}
        >
        </Slider>
        </Box>
        </Accordion>
        <Accordion title='Жанры' defaultOpen>
          <CheckboxList name='genres' values={Genres} register={register} />
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
