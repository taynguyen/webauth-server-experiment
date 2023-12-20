package pwl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) BeginLogin(c *gin.Context) {
	// Example, we will create a new user
	user := WAUser{
		WaID:          randomUserID(),
		WaName:        "Username",
		WaDisplayName: "Displayname",
	}

	options, session, err := h.webAuthn.BeginLogin(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// store the session values
	setUserSession(user.WaID, session)
	setUser(user.WaID, user)

	c.JSON(http.StatusOK, options) // return the options generated
	// options.publicKey contain our registration options
}

func (h Handler) FinishLogin(c *gin.Context) {
	// user := datastore.GetUser() // Get the user

	// // Get the session data stored from the function above
	// session := datastore.GetSession()

	// credential, err := webAuthn.FinishLogin(user, session, r)
	// if err != nil {
	// 	// Handle Error and return.

	// 	return
	// }

	// // Handle credential.Authenticator.CloneWarning

	// // If login was successful, update the credential object
	// // Pseudocode to update the user credential.
	// user.UpdateCredential(credential)
	// datastore.SaveUser(user)

	c.JSON(http.StatusOK, "Login Success")
}
