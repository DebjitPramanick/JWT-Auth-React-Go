import React from 'react'

const RegisterForm = () => {
  return (
    <div className='form-container'>
        <form>
            <input placeholder='Enter name'></input>
            <input placeholder='Enter email'></input>
            <input placeholder='Enter password'></input>
            <button>Register</button>
        </form>
    </div>
  )
}

export default RegisterForm