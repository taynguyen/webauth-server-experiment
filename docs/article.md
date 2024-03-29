# Web-auth

## Fundamentals
### Why ?
In today's digital world, where data breaches and password fatigue run rampant, a robust and secure authentication system is paramount. Enter web-auth, a game-changer in online security that offers a powerful alternative to the vulnerable username-password combo.

#### Beyond Security: Web-auth brings more than just ironclad security.

 - **Convenience**: Say goodbye to password resets and forgotten credentials. Logins become faster and more intuitive, using familiar methods like fingerprints or facial recognition.
 - **Privacy**: No more storing passwords on vulnerable servers. Private keys reside solely on your trusted device, offering unparalleled data protection.
 - **Scalability**: Web-auth seamlessly integrates with existing systems, making it easy for websites to adopt and users to embrace.


### What is it ?
Quick summary what is web-auth
E.g: Web-auth, formally known as Web Authentication API (WebAuthn), is a revolutionary web standard for user authentication. It leverages public-key cryptography, where users have a private key (kept secret) and a public key (shared with the website). During login, instead of entering a password, users "present" their private key on a trusted device (phone, security key) – proving their identity without revealing any sensitive information.

### How it works ?
Example:

Imagine web-auth as a secure handshake between you and the website.

 - **Registration**: You create a private-public key pair on your device. The public key is sent to the website, while the private key remains safely locked away.
 - **Login**: When prompted, you use your device to "present" your private key (fingerprint scan, PIN, etc.).
 - **Verification**: The website sends a challenge to your device. Your device signs the challenge using your private key, proving you possess it.
 - **Authentication**: The website verifies the signed challenge, confirming your identity without ever learning your private key.

This "proof-of-possession" eliminates the need for passwords, significantly boosting security. Hackers can't steal what they can't see, making phishing and brute-force attacks virtually obsolete.



## Implementations in golang
 ### Introduce the library we're using
 The library we're using is [webauthn](https://github.com/go-webauthn/webauthn), which is a Go implementation that provides a passwordless solution for users.
 ### How we use it
   Note: Show code, explain it. Also take a look at the implementation code.
 #### Register
 This is the example code in the library:
 ```
 package example

func BeginRegistration(w http.ResponseWriter, r *http.Request) {
	user := datastore.GetUser() // Find or create the new user  
	options, session, err := webAuthn.BeginRegistration(user)
	// handle errors if present
	// store the sessionData values 
	JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options
}

func FinishRegistration(w http.ResponseWriter, r *http.Request) {
	user := datastore.GetUser() // Get the user
	
	// Get the session data stored from the function above
	session := datastore.GetSession()
		
	credential, err := webAuthn.FinishRegistration(user, session, r)
	if err != nil {
		// Handle Error and return.

		return
	}
	
	// If creation was successful, store the credential object
	// Pseudocode to add the user credential.
	user.AddCredential(credential)
	datastore.SaveUser(user)

	JSONResponse(w, "Registration Success", http.StatusOK) // Handle next steps
}
```
The registration process begins by querying for the user in the database. Once the user is retrieved, we call the `webAuthn.BeginRegistration(user)` function from the library to generate registration data (the `session`). This `session` value should be stored and retrieved later on. The `options` should be returned in the response.

To finish the registration, first both the user and the session is retrieved from the database. After that, we call the `webAuthn.FinishRegistration(user, session, r)` function from the library to get the `credential` for the user. We then add this `credential` to the user struct, and store the user back to the database.

This process is implemented in `/go-api/pkg/handler/v1/pwl/registration.go`.

 #### Login
 This is the example code in the library:
  ```
 package example

func BeginLogin(w http.ResponseWriter, r *http.Request) {
	user := datastore.GetUser() // Find the user
	
	options, session, err := webAuthn.BeginLogin(user)
	if err != nil {
		// Handle Error and return.

		return
	}
	
	// store the session values
	datastore.SaveSession(session)
	
	JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options
}

func FinishLogin(w http.ResponseWriter, r *http.Request) {
	user := datastore.GetUser() // Get the user 
	
	// Get the session data stored from the function above
	session := datastore.GetSession()
	
	credential, err := webAuthn.FinishLogin(user, session, r)
	if err != nil {
		// Handle Error and return.

		return
	}

	// Handle credential.Authenticator.CloneWarning

	// If login was successful, update the credential object
	// Pseudocode to update the user credential.
	user.UpdateCredential(credential)
	datastore.SaveUser(user)
	
	JSONResponse(w, "Login Success", http.StatusOK)
}
```
Users should log in through `webAuthn.BeginLogin(user)` to get the `session`. This `session` will then be stored, and the `options` is returned to the caller.

To finish logging in, we call `webAuthn.FinishLogin(user, session, r)` and then update the user's credential accordingly.

This process is implemented in `/go-api/pkg/handler/v1/pwl/login.go`.

## References
1. https://github.com/go-webauthn/webauthn

## Conclusion
Web-auth is an undeniable leap forward in online security, paving the way for a more secure, convenient, and private web experience. Its robust cryptographic foundation and user-friendly features make it the future of online authentication, offering peace of mind to both users and website owners. So, the next time you log in, look for the web-auth option – it's time to ditch the passwords and embrace a secure digital future.
