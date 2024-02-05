import { useState } from 'react'


import TodoList from './page/toDoList'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <TodoList></TodoList>
    </>
  )
}

export default App
