import ReactPaginate from 'react-paginate'

import ArrowIcon from '../icons/ArrowIcon'

import './style.scss'

const Pagination = (props) => {
  // if (props.pageCount <= 1) return null
  return (
    <div className='pagination'>
      <ReactPaginate
        className='pagination__body'
        nextLabel={<ArrowIcon />}
        previousLabel={<ArrowIcon />}
        renderOnZeroPageCount={null}
        {...props}
      />
    </div>
  )
}

export default Pagination
