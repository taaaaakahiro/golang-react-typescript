import type { NextPage } from 'next'
import React, {useEffect, useState} from 'react'
import {Parent} from '../components/SampleComponent1'

type test = {
  st : string
  num : number
  big : bigint // 2**53以上
  bool : boolean
  array_s : string[]
  array_n : readonly number[]
  array_sn : ReadonlyArray<(number | string)[]>
  an : any
  unk : unknown
}

let test_2 : {
  name : string
  code : number
  array : string[]
} = {
  name : 'test',
  code : 89,
  array : ['1', '2', '3']
}



function log() : void {
  console.log('run log method')
  console.log(test_2.name) // test
  console.log(test_2.code) // undefined
  test_2.array.map(( _, i ) => { // 1, 2, 3
    console.log(_)
  })
} 



const Home: NextPage = () => {
  const [id, setId] =useState(0)

  useEffect(() => {
    fetch("http://localhost:8080/rest", {method: 'GET'})
    .then((res)=> {
      setId(res.status)
    })
    log()
  }, [])

  let i = 6666

  return (
  <>
    <Parent></Parent>{'/' + i}
    <h4>
      HTTP Response Code -&gt; {id === 200 ? id : '失敗'}
    </h4>
  </>
  )
}

export default Home
