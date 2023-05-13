import { Route, Routes } from 'react-router-dom'

import Footer from './components/Footer'
import Header from './components/Header'
import LoginPage from './pages/LoginPage'
import MainPage from './pages/MainPage'
import RegistrationPage from './pages/RegistrationPage'
import AuthService from './services/authService'
import { useEffect, useState } from 'react';

import './scss/app.scss'
import ProductService from './services/productService'
import StaffService from './services/staffService'
import CatalogPage from './pages/CatalogPage'
import ProductPage from './pages/ProductPage'
import OrderPage from './pages/Orders'
import FavouritesPage from './pages/Favourites'
import CartPage from './pages/CartPage'

const App = () => {
  const [user, setUser] = useState(null)
  let User = {
    Email: "evi13@tpu.ru",
    Pass: "159753Kupa"
  }
  useEffect(() => {
    if (JSON.parse(localStorage.getItem('cart')) === null ){
      localStorage.setItem('cart', JSON.stringify([]));
    }
    AuthService.loginCookie().then(
      (res) =>{
        setUser(res)
      }
    )
  }, [])
  let RegUser = {
    Name: "Pipa",
    Email: "2002egor@gmail.com",
    Pass: "159753Kupa"
  }
  let data = {
    Code: "yztdnt",
    Email: "2002egor@gmail.com",
  }
  let fav = {
    User: 28,
    Book: 32
  }
  let SearchOpt = {
    Str:        "Азимов",
    Genre:      "",
    Publishers: "",
    StartDate:  "",
    EndDate: ""
  }
  let Order = {
    UserID:   28,
    Adress:   "Pushk",
    Status:   "Cherno",
    Book_IDs: [32],
    Amounts:  [1]
  }
  let Book = {
    ID:         33,
		Name:       "Город и звезды",
		Price:      "555.55",
		Discount:   0,
		ISBN:       "978-5-17-105787-9",
		Photo:      "default",
		Desc:       "Артура Кларка Айзека Азимова и Роберта Хайнлайна называют",
		Dimensions: [18, 11, 16],
		Authors:    ["Артур Чарльз Кларк", "Айзек Азимов"],
		Publishers: ["АСТ"],
		Genres:     ["Фантастика"],
		Amount:     5,
		Date:       "2021-10-5",
  }
   return (
    <div className='app'>
      <Header user = {user} setUser = {setUser}  />
      <main className='main'>
        <Routes>
          <Route path='/login' element={<LoginPage user = {user} setUser = {setUser}/>} />
          <Route path='' element={<MainPage user = {user} setUser = {setUser}/>} />
          <Route path='/registration' element={<RegistrationPage setUser = {setUser}/>} />
          <Route path='/catalog' element={<CatalogPage />} />
          <Route path='/book/:id' element={<ProductPage />} />
          <Route path='/favourites' element={<FavouritesPage />} />
          <Route path='/orders' element={<OrderPage />} />
          <Route path='/cart' element={<CartPage />} />
        </Routes>
      </main>
      <Footer />
    </div>
  ) 
 // 
}

export default App
