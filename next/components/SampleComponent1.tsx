import React from "react"

type Props = {
  text: string
  num : number
  children: React.ReactNode
}

const SampleComponent5: React.FC<Props>= (props) => {
  return (
    <div>
      <h1>Hello {props.text}!</h1>
      <p>{props.children}</p>
      <p>{props.num}</p>
    </div>
  )
}


/* ---------- 呼び出す側 ---------- */
export const Parent: React.FC = () => {
  let i : number = 111
  let s : string = "TypeScript & Golang"
  let s_array : string[] = ['test1', 'test2']
  let n_array : number[] = [1111, 2222]
  let c_array : [string, number] = ['test3', 2222]
  let any_array = ['test4', 3333] // any
  return (
    <>
      <SampleComponent5 text={s} num={i}>
        Props.childrenを表示
      </SampleComponent5>
        <div>
          {n_array[0]}
        </div>
        <div>
          {s_array[0]}
        </div>
        <div>
          {c_array}
        </div>
      
      
    </>
  )
}