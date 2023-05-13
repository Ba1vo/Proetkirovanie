import { useEffect, useState } from 'react'

import classNames from 'classnames'
import { useDispatch, useSelector } from 'react-redux'
import { useMediaQuery } from 'react-responsive'
import { useNavigate, useSearchParams } from 'react-router-dom'

import { getOrdersProducts } from '../../slices/orderProductSlice'
import ErrorIndicator from '../ErrorIndicator'
import Spinner from '../Spinner'
import { useFocusEffect } from '@react-navigation/native';

import './style.scss'
import OrdersListView from './OrderListView'

const OrdersList = () => {
  const [searchParams, setSearchParams] = useSearchParams()
  const dispatch = useDispatch()

  const { products, count, loading, error } = useSelector(
    (state) => state.orders
  )

  const isMobile = useMediaQuery({ query: '(max-width: 1080px)' })

  useEffect(() => {
    dispatch(
      getOrdersProducts()
    )
  }, [])

  const currentPage = searchParams.get('page')
    ? searchParams.get('page') - 1
    : 0

  return (
    <div
      className={classNames('order-list-container', {
        'order-list-container_not-loaded': loading || error,
      })}
    >
      {loading ? <Spinner /> : null}
      {error ? <ErrorIndicator /> : null}
      {!(loading || error) ? (
        <OrdersListView
          products={products}
          forcePage={currentPage}
          count = {count}
        />
      ) : null}
    </div>
  )
}

export default OrdersList
