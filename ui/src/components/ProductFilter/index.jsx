import { useEffect } from 'react'

import classNames from 'classnames'
import PropTypes from 'prop-types'
import { useForm } from 'react-hook-form'
import { useDispatch, useSelector } from 'react-redux'
import { useSearchParams } from 'react-router-dom'

import ProductFilterView from './ProductFilterView'
import { objectToQueryString, queryStringToObject } from './utils'
import { getFilters } from '../../slices/catalogProductsSlice'
import ErrorIndicator from '../ErrorIndicator'
import Spinner from '../Spinner'

import './style.scss'

const ProductFilter = () => {
  const { register, handleSubmit, reset, setValue, watch } = useForm()

  const [searchParams, setSearchParams] = useSearchParams()

  const { filters, error, loading } = useSelector(
    (state) => state.catalogProducts.filters
  )
  const dispatch = useDispatch()

  const submitHandler = (data) => {

    console.log(data)
    const { MinPrice, MaxPrice } = filters

    if (data['min_price'] && Number(data['min_price']) < MinPrice) {
      data['min_price'] = MinPrice
      setValue('min_price', MinPrice)
    }

    if (data['max_price'] && Number(data['max_price']) > MaxPrice) {
      data['max_price'] = MaxPrice
      setValue('max_price', MaxPrice)
    }

    if (
      Number(data['min_price']) &&
      Number(data['max_price']) &&
      Number(data['min_price']) > Number(data['max_price'])
    ) {
      const temp = data['min_price']
      data['min_price'] = data['max_price']
      setValue('min_price', data['max_price'])
      data['max_price'] = temp
      setValue('max_price', temp)
    }

    data.page = 1
    setSearchParams(objectToQueryString(data))

    window.scrollTo({ top: 174, behavior: 'smooth' })
  }

  const resetHandler = () => {
    const order = searchParams.get('order_by')
    setSearchParams(order ? `order_by=${order}` : '')
    reset()

    window.scrollTo({ top: 174, behavior: 'smooth' })
  }

  useEffect(() => {
    const filtersFromParams = queryStringToObject(searchParams)
    if (Object.keys(filtersFromParams).length === 0) {
      reset()
    }
    for (const key in filtersFromParams) {
      setValue(key, filtersFromParams[key])
    }
  }, [searchParams])

  useEffect(() => {
    dispatch(getFilters())
  }, [])

  return (
    <div
      className={classNames('product-filter-container', {
        'product-filter-container_not-loaded': loading || error,
      })}
    >
      {loading ? <Spinner className /> : null}
      {error ? <ErrorIndicator /> : null}
      {!(loading || error) ? (
        <ProductFilterView
          filters={filters}
          register={register}
          watch={watch}
          setValue={setValue}
          onSubmit={handleSubmit(submitHandler)}
          onReset={resetHandler}
        />
      ) : null}
    </div>
  )
}

ProductFilter.propTypes = {
  className: PropTypes.string,
}

export default ProductFilter
