import axios from 'axios';
import React, { useEffect, useState } from 'react';


export default function About() {
    const [about, setAbout] = useState('')

    useEffect(() => {
        const call = async() => {
            const response = await axios.get('http://127.0.0.1:8000/about')
            console.log(response);
            setAbout(response.data)
        }

        call()

    }, [])
    


    return (
        <div className="row mt-5">
            <div className="col-12 order-lg-1">
                <h3 className="mb-4">About the Go Music Store</h3>
                <div>{about}</div>
                
            </div>
        </div>);
}