# Web-auth

## Fundamentals
### Why ?
In today's digital world, where data breaches and password fatigue run rampant, a robust and secure authentication system is paramount. Enter web-auth, a game-changer in online security that offers a powerful alternative to the vulnerable username-password combo.

Beyond Security: Web-auth brings more than just ironclad security.

Convenience: Say goodbye to password resets and forgotten credentials. Logins become faster and more intuitive, using familiar methods like fingerprints or facial recognition.
Privacy: No more storing passwords on vulnerable servers. Private keys reside solely on your trusted device, offering unparalleled data protection.
Scalability: Web-auth seamlessly integrates with existing systems, making it easy for websites to adopt and users to embrace.


### What is it ?
Quick summary what is web-auth
E.g: Web-auth, formally known as Web Authentication API (WebAuthn), is a revolutionary web standard for user authentication. It leverages public-key cryptography, where users have a private key (kept secret) and a public key (shared with the website). During login, instead of entering a password, users "present" their private key on a trusted device (phone, security key) – proving their identity without revealing any sensitive information.

### How it works ?
Example:
Imagine web-auth as a secure handshake between you and the website.

Registration: You create a private-public key pair on your device. The public key is sent to the website, while the private key remains safely locked away.
Login: When prompted, you use your device to "present" your private key (fingerprint scan, PIN, etc.).
Verification: The website sends a challenge to your device. Your device signs the challenge using your private key, proving you possess it.
Authentication: The website verifies the signed challenge, confirming your identity without ever learning your private key.

This "proof-of-possession" eliminates the need for passwords, significantly boosting security. Hackers can't steal what they can't see, making phishing and brute-force attacks virtually obsolete.



## Implementations in golang
 ### Introduce the library we're using
 ### How we use it
   Show code, explain it. Also take a look at the implementation code.

## References

## Conclusion
Web-auth is an undeniable leap forward in online security, paving the way for a more secure, convenient, and private web experience. Its robust cryptographic foundation and user-friendly features make it the future of online authentication, offering peace of mind to both users and website owners. So, the next time you log in, look for the web-auth option – it's time to ditch the passwords and embrace a secure digital future.
