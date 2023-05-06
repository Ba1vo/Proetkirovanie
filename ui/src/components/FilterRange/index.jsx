import classNames from 'classnames'
import PropTypes from 'prop-types'

import { NumberField } from '../form-inputs'

import './style.scss'

const FilterRange = ({
  min,
  max,
  minInputName,
  maxInputName,
  setValue,
  register,
  className,
}) => {
  return (
    <div className={classNames('filter-range', className)}>
      <NumberField
        className='filter-range__field'
        name={minInputName}
        setValue={setValue}
        register={register}
        placeholder={`От ${min}`}
        small
      />
      <div className='filter-range__dash' />
      <NumberField
        className='filter-range__field'
        name={maxInputName}
        setValue={setValue}
        register={register}
        placeholder={`До ${max}`}
        small
      />
    </div>
  )
}

FilterRange.propTypes = {
  className: PropTypes.string,
}

export default FilterRange
