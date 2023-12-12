package pwl

import (
	"github.com/gin-gonic/gin"
)

func (h Handler) BeginRegistration(c *gin.Context) {
	// user := datastore.GetUser() // Find or create the new user
	// options, session, err := h.webAuthn.BeginRegistration(user)
	// // handle errors if present
	// // store the sessionData values
	// JSONResponse(w, options, http.StatusOK) // return the options generated
	// // options.publicKey contain our registration options
}

func (h Handler) FinishRegistration(c *gin.Context) {
	// user := datastore.GetUser() // Get the user

	// // Get the session data stored from the function above
	// session := datastore.GetSession()

	// credential, err := h.webAuthn.FinishRegistration(user, session, r)
	// if err != nil {
	// 	// Handle Error and return.

	// 	return
	// }

	// // If creation was successful, store the credential object
	// // Pseudocode to add the user credential.
	// user.AddCredential(credential)
	// datastore.SaveUser(user)

	// JSONResponse(w, "Registration Success", http.StatusOK) // Handle next steps
}
