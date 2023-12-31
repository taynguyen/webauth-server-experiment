import React from 'react';
import logo from './logo.svg';
import './App.css';
import axios from 'axios';
import { startAuthentication, startRegistration } from '@simplewebauthn/browser';

function App() {
  // const [authenticator, setAuthenticator] = React.useState<any>(null)
  
  const onRegisterBegin = async () => {
    const res = await axios.post('http://localhost:3000/api/v1/pwl/user/registration/begin')
    const resData = res.data
    
    try {
      // console.log('registerBegin: ', {
      //   challenge: registerBegin.publicKey.challenge,
      //   rpId: registerBegin.publicKey.rp.id,
      //   timeout: registerBegin.publicKey.timeout,
      //   userVerification: registerBegin.publicKey.authenticatorSelection.userVerification,
      // })
      const authRes = await startRegistration(resData.publicKey)

      const verifyRes = await axios.post('http://localhost:3000/api/v1/pwl/user/registration/finish', authRes, {
        headers: {
          'user-id': resData.publicKey.user.id,
        }
      })
      const verifyData = verifyRes.data
      console.log('verifyData: ', verifyData)
      
    } catch (err) {
      console.error(err)
    }


  }
  const onRegisterFinish = () => {
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

      <button onClick={onRegisterBegin}>Register Begin</button>
      <button onClick={onRegisterFinish}>Register Finish</button>
    </div>
  );
}

export default App;
