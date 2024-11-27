import React from 'react';
import _, { useState, useEffect } from 'react';


interface IUser {
  id: number
  email: string
  name: string

}

function App() {
  const [users, setUsers] = useState<IUser[]>([])

  const getUsersHandler = async () => {
    const response = await fetch('http://localhost:3000/api/users')
    const data: IUser[] = await response.json()
    console.log(data)
    setUsers(data)
  }

  const postUserHandler = async (e: React.FormEvent) => {
    e.preventDefault();
  
    const form = e.currentTarget as HTMLFormElement;
    const formData = new FormData(form);
  
    const data: any = {};
    formData.forEach((value, key) => {
      data[key] = value;
    });
  

    console.log(data)
    const response = await fetch('http://localhost:3000/api/users/create', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    });
  
    const responseData = await response.json();
    console.log(responseData);
  };

  // const incrementCount = async () => {
  //   const response = await fetch('http://localhost:3000/api/counter/increment', {
  //     method: 'POST',
  //   });
  //   const data = await response.json();
  //   setCount(data.count);
  // };

  // const resetCount = async () => {
  //   const response = await fetch('http://localhost:3000/api/counter/reset', {
  //     method: 'POST',
  //   });
  //   const data = await response.json();
  //   setCount(data.count);
  // };

  


  return (
    <div style={{ textAlign: 'center', marginTop: '50px' }}>
      <button onClick={getUsersHandler} style={{ margin: '10px', padding: '10px' }}>
        Получить пользователей
      </button>




      <form onSubmit={postUserHandler}>

        <input placeholder='name' name='name' required />

        <input placeholder='email' name='email' required />


        <button type='submit' style={{ margin: '10px', padding: '10px' }}>
          Создать пользователя
        </button>

      </form>



      <div className='users'>


        {users?.map(user => (
          <div key={user.id}>
            <h2>{user.name}</h2>
            <p>{user.id}</p>
            <p>{user.email}</p>
          </div>
        ))}


      </div>



    </div>
  );
}

export default App;
