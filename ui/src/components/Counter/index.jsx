import { useState } from 'react'

import MinusIcon from '../icons/MinusIcon'
import PlusIcon from '../icons/PlusIcon'

import './style.scss'

const Counter = () => {
  const [count, setCount] = useState(1)

  const handleDec = () => {
    if (count > 1) setCount((count) => count - 1)
  }

  const handleInc = () => {
    setCount((count) => count + 1)
  }

  return (
    <div className='counter'>
      <div className='counter__minus-div' onClick={handleDec}>
        <MinusIcon className='counter__minus' />
      </div>
      <div className='counter__value'>{count}</div>
      <div className='counter__plus-div' onClick={handleInc}>
        <PlusIcon className='counter__plus' />
      </div>
    </div>
  )
}

export default Counter
