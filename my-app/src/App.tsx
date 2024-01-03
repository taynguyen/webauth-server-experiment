import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import { startAuthentication, startRegistration } from '@simplewebauthn/browser';

function App() {
  // const [authenticator, setAuthenticator] = React.useState<any>(null)
  const [userId, setUserId] = React.useState<string>('')
  
  const onRegisterBegin = async () => {
    const res = await axios.post('http://localhost:3000/api/v1/pwl/user/registration/begin')
    const resData = res.data
    
    try {
      const authRes = await startRegistration(resData.publicKey)

      const verifyRes = await axios.post('http://localhost:3000/api/v1/pwl/user/registration/finish', authRes, {
        headers: {
          'user-id': resData.publicKey.user.id,
        }
      })
      const verifyData = verifyRes.data
      console.log('verifyData: ', verifyData)

      setUserId(resData.publicKey.user.id)
      
    } catch (err) {
      console.error(err)
    }
  }

  const onLogin = async () => {
    const res = await axios.post('http://localhost:3000/api/v1/pwl/user/login/begin', {}, 
    {
      headers: {
        'user-id': userId,
      }
    })

    const resData = res.data
    console.log('resData: ', resData)
    const authRes = await startAuthentication(resData.publicKey)


    // Validate login
    const verifyRes = await axios.post('http://localhost:3000/api/v1/pwl/user/login/finish', authRes,
    {
      headers: {
        'user-id': userId,
      }
    })
    console.log('verifyRes: ', verifyRes)
  }

  return (
    <div className="App">
      {/* <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header> */}

      <button onClick={onRegisterBegin}>Register</button>
      <button onClick={onLogin}>Login</button>
    </div>
  );
}

export default App;
