import React, { useState } from 'react'
import { login } from '../api/api.helper'

const LoginForm = () => {

  const [data, setData] = useState({
    email: '',
    password: ''
  })

  const handleLogin = async (e: any) => {
    e.preventDefault()
    try {
      await login(data)
    } catch (err) {
      alert(err);
    }
  }

  return (
    <div className='form-container'>
        <form onSubmit={handleLogin}>
            <input placeholder='Enter email' type={'email'} value={data.email} onChange={(e) => setData({...data, email: e.target.value})}></input>
            <input placeholder='Enter password' type={'password'} value={data.password} onChange={(e) => setData({...data, password: e.target.value})}></input>
            <button type='submit'>Login</button>
        </form>
    </div>
  )
}

export default LoginForm