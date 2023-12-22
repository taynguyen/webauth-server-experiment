package pwl

import (
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) BeginRegistration(c *gin.Context) {
	// Example, we will create a new user

	user := WAUser{
		WaID:          randomUserID(),
		WaName:        "WaName",
		WaDisplayName: "WaDisplayName",
	}

	options, session, err := h.webAuthn.BeginRegistration(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	// handle errors if present
	// store the sessionData values
	// fmt.Printf("session: %+v\n", session)
	setUserSession(user.WaID, session)
	setUser(user.WaID, user)
	fmt.Println("userID: ", b64Encoding.EncodeToString(user.WaID))

	// JSONResponse(w, options, http.StatusOK) // return the options generated
	// options.publicKey contain our registration options
	c.JSON(http.StatusOK, options)

}

func randomUserID() []byte {
	buf := make([]byte, 64)
	rand.Read(buf)
	return buf
}

func (h Handler) FinishRegistration(c *gin.Context) {
	encodedUserID := c.GetHeader("user-id")
	fmt.Println("encodedUserID: ", encodedUserID)
	// user := datastore.GetUser() // Get the user
	user, existed := userRepo[encodedUserID]
	if !existed {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	// Get the session data stored from the function above
	// session := datastore.GetSession()
	session, existed := userSessionRepo[encodedUserID]
	if !existed {
		c.JSON(http.StatusBadRequest, "session not found")
		return
	}
	if session == nil {
		c.JSON(http.StatusBadRequest, "session is nil")
		return
	}

	credential, err := h.webAuthn.FinishRegistration(user, *session, c.Request)
	if err != nil {
		// Handle Error and return.
		c.JSON(http.StatusBadRequest, err)
		return
	}
	fmt.Printf("credential: %+v\n", credential)
	if credential == nil {
		c.JSON(http.StatusBadRequest, "credential is nil")
		return
	}

	// If creation was successful, store the credential object
	// Pseudocode to add the user credential.
	// user.AddCredential(credential)
	// datastore.SaveUser(user)
	user.WaCredentials = append(user.WaCredentials, *credential)
	setUser(user.WaID, user)

	// JSONResponse(w, "Registration Success", http.StatusOK) // Handle next steps
	c.JSON(http.StatusOK, "Registration Success")
}
