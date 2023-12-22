package pwl

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h Handler) BeginLogin(c *gin.Context) {
	encodedUserID := c.GetHeader("user-id")
	user, existed := userRepo[encodedUserID]
	if !existed {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	options, session, err := h.webAuthn.BeginLogin(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// store the session values
	setUserSession(user.WaID, session)

	c.JSON(http.StatusOK, options) // return the options generated
}

func (h Handler) FinishLogin(c *gin.Context) {
	encodedUserID := c.GetHeader("user-id")
	user, existed := userRepo[encodedUserID]
	if !existed {
		c.JSON(http.StatusBadRequest, "user not found")
		return
	}

	// // Get the session data stored from the function above
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

	credential, err := h.webAuthn.FinishLogin(user, *session, c.Request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}

	// If login was successful, update the credential object
	// Pseudocode to update the user credential.
	// user.AddCredential(credential)
	// datastore.SaveUser(user)
	user.WaCredentials = append(user.WaCredentials, *credential)
	setUser(user.WaID, user)

	c.JSON(http.StatusOK, "Login Success")
}
