import { useState } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'

import ArrowIcon from '../icons/ArrowIcon'

import './style.scss'

const Accordion = ({ title, defaultOpen = false, children, className }) => {
  const [open, setOpen] = useState(defaultOpen)

  const toggleAccordion = () => {
    setOpen((open) => !open)
  }

  return (
    <div className={classNames('accordion', className)}>
      <div className='accordion__title' onClick={toggleAccordion}>
        {title}
        <ArrowIcon
          className={classNames('accordion__icon', {
            'accordion__icon_open': open,
          })}
        />
      </div>
      <div
        className={classNames('accordion__content', {
          'accordion__content_open': open,
        })}
      >
        <div className='accordion__content-container'>{children}</div>
      </div>
    </div>
  )
}

Accordion.propTypes = {
  title: PropTypes.string,
  children: PropTypes.node,
  className: PropTypes.string,
}

export default Accordion
