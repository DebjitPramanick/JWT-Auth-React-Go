import React, { useEffect, useState } from 'react'
import { getUser } from '../api/api.helper'

const Home = () => {

    const [user, setUser] = useState<any | null>(null)

    useEffect(() => {
        fetchUser()
    }, [])


    const fetchUser = async () => {
        try {
            const res: any = await getUser()
            setUser(res.data)
        } catch (err) {
            alert(err)
        }
    }

    return (
        <div>
            <h1>{user?.name}</h1>
        </div>
    )
}

export default Home