import type { NextPage } from 'next'
import React, {useEffect, useState} from 'react'
import {Parent} from '../components/SampleComponent1'




const Home: NextPage = () => {
  const [id, setId] =useState('test')

  useEffect(() => {
    fetch("http://localhost:8080/rest", {method: 'GET'}) //CORS error
    .then((res)=> console.log(res.status))
    
  }, [])

  let i = 6666

  return (
  <>
    <Parent></Parent>{i}
  </>
  )
}

export default Home
