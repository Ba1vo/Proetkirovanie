import classNames from 'classnames'
import PropTypes from 'prop-types'

import CheckboxList from '../CheckboxList'
import { TextField } from '../form-inputs'
import LoupeIcon from '../icons/LoupeIcon'

import './style.scss'

const CheckboxSearchList = ({
  name,
  values,
  searchInputName,
  watch,
  register,
  className,
}) => {
  const term = watch(searchInputName) ?? ''
  const visibleValues = values?.filter((value) =>
    value.toLowerCase().includes(term.trim().toLowerCase())
  )

  return (
    <div className={classNames('checkbox-search-list', className)}>
      <TextField
        className='checkbox-search-list__field'
        name={searchInputName}
        register={register}
        placeholder='Поиск'
        small
      >
        <LoupeIcon className='checkbox-search-list__search-icon' />
      </TextField>
      {visibleValues?.length === 0 ? (
        <span className='checkbox-search-list__not-found-message'>
          Не найдено
        </span>
      ) : (
        <CheckboxList
          className='checkbox-search-list__list'
          name={name}
          values={visibleValues}
          register={register}
        />
      )}
    </div>
  )
}

CheckboxSearchList.propTypes = {
  className: PropTypes.string,
}

export default CheckboxSearchList
