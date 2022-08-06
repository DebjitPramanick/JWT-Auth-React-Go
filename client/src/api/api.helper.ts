import axios from "axios"

const API = 'http://127.0.0.1:4000/api'

const api = axios.create({
    baseURL: API
})

export const register = async (data: any) => {
    const res = await api.post('/register', data)
    .then(res => res)
    .catch(err => {
        console.log("ERROR Register ==> ", err)
        throw new Error("Could not register user.")
    })
    return res;
}

export const login = async (data: any) => {
    const res = await api.post('/login', data)
    .then(res => res)
    .catch(err => {
        console.log("ERROR Login ==> ", err)
        throw new Error("Could not login user.")
    })
    return res;
}

export const logout = async () => {
    const res = await api.post('/logout')
    .then(res => res)
    .catch(err => {
        console.log("ERROR Logout ==> ", err)
        throw new Error("Could not logout user.")
    })
    return res;
}

export const getUser = async (data: any) => {
    const res = await api.get('/user')
    .then(res => res)
    .catch(err => {
        console.log("ERROR Fetching user ==> ", err)
        throw new Error("Could not fetch user data.")
    })
    return res;
}