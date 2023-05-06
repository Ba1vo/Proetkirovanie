import classNames from 'classnames'
import PropTypes from 'prop-types'

import { Checkbox } from '../form-inputs'

import './style.scss'

const CheckboxList = ({ name, values, register, className }) => {
  return (
    <div className={classNames('checkbox-list', className)}>
      {values?.map((value) => (
        <Checkbox
          className='checkbox-list__checkbox'
          key={value}
          name={name}
          label={value}
          value={value}
          register={register}
        />
      ))}
    </div>
  )
}

CheckboxList.propTypes = {
  className: PropTypes.string,
}

export default CheckboxList
