import { useState } from 'react'

import classNames from 'classnames'

import ArrowIcon from '../../icons/ArrowIcon'

import './style.scss'

const Select = ({ options, selectedOption, onChange, className }) => {
  const [isOpen, setIsOpen] = useState(false)

  const onSelectClick = () => {
    setIsOpen((prevValue) => !prevValue)
  }

  const onSelectBlur = () => {
    setIsOpen(false)
  }

  const onOptionSelect = (option) => {
    if (option !== selectedOption) {
      onChange(option)
    }
  }

  return (
    <div
      className={classNames('select', className)}
      onClick={onSelectClick}
      onBlur={onSelectBlur}
      tabIndex={0}
    >
      <span className='select__value'>{selectedOption.label}</span>
      <ArrowIcon
        className={classNames('select__icon', {
          'select__icon_open': isOpen,
        })}
      />
      <ul
        className={classNames('select__option-list', {
          'select__option-list_open': isOpen,
        })}
      >
        {options.map((option) => (
          <li
            key={option.value}
            className={classNames('select__option', {
              'select__option_selected': option === selectedOption,
            })}
            onClick={() => onOptionSelect(option)}
          >
            {option.label}
          </li>
        ))}
      </ul>
    </div>
  )
}

export default Select
