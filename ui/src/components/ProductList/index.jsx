import { useEffect, useState } from 'react'

import classNames from 'classnames'
import { useDispatch, useSelector } from 'react-redux'
import { useMediaQuery } from 'react-responsive'
import { useNavigate, useSearchParams } from 'react-router-dom'

import ProductListView from './ProductListView'
import { getProducts } from '../../slices/catalogProductsSlice'
import ErrorIndicator from '../ErrorIndicator'
import Spinner from '../Spinner'
import { useFocusEffect } from '@react-navigation/native';

import './style.scss'

const ProductList = () => {
  const [searchParams, setSearchParams] = useSearchParams()
  const dispatch = useDispatch()

  const { products, count, loading, error } = useSelector(
    (state) => state.catalogProducts.products
  )
  const isMobile = useMediaQuery({ query: '(max-width: 1080px)' })

  const productsPerPage = isMobile ? 8 : 9
  const pageCount = Math.ceil(count / productsPerPage)

  const onPageClick = (event) => {
    setSearchParams((searchParams) => {
      searchParams.set('page', event.selected + 1)
      return searchParams
    })
   
    window.scrollTo({ top: 174, behavior: 'smooth' })
  }


  useEffect(() => {
    dispatch(
      getProducts({
        queryString: searchParams.toString(),
        pageSize: productsPerPage,
      })
    )
  }, [searchParams, productsPerPage])

  const currentPage = searchParams.get('page')
    ? searchParams.get('page') - 1
    : 0

  return (
    <div
      className={classNames('product-list-container', {
        'product-list-container_not-loaded': loading || error,
      })}
    >
      {loading ? <Spinner /> : null}
      {error ? <ErrorIndicator /> : null}
      {!(loading || error) ? (
        <ProductListView
          products={products}
          pageCount={pageCount}
          forcePage={currentPage}
          onPageClick={onPageClick}
        />
      ) : null}
    </div>
  )
}

export default ProductList
