import React, { useState } from 'react'
import { register } from '../api/api.helper'

const RegisterForm = () => {

  const [data, setData] = useState({
    name: '',
    email: '',
    password: ''
  })

  const handleRegister = async (e: any) => {
    e.preventDefault()
    try {
      await register(data)
    } catch (err) {
      alert(err);
    }
  }

  return (
    <div className='form-container'>
      <form onSubmit={handleRegister}>
        <input placeholder='Enter name' type={'text'} value={data.name} onChange={(e) => setData({...data, name: e.target.value})}></input>
        <input placeholder='Enter email' type={'email'} value={data.email} onChange={(e) => setData({...data, email: e.target.value})}></input>
        <input placeholder='Enter password' type={'password'} value={data.password} onChange={(e) => setData({...data, password: e.target.value})}></input>
        <button type='submit'>Register</button>
      </form>
    </div>
  )
}

export default RegisterForm