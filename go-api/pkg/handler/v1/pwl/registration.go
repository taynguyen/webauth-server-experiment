package pwl

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-webauthn/webauthn/webauthn"
)

type WAUser struct {
	WaID          []byte
	WaDisplayName string
	WaName        string
	WaCredentials []webauthn.Credential
	WaIcon        string
}

func (u WAUser) WebAuthnID() []byte {
	return u.WaID
}

func (u WAUser) WebAuthnDisplayName() string {
	return u.WaDisplayName
}

func (u WAUser) WebAuthnName() string {
	return u.WaName
}

func (u WAUser) WebAuthnCredentials() []webauthn.Credential {
	return u.WaCredentials
}

func (u WAUser) WebAuthnIcon() string {
	return u.WaIcon
}

var b64Encoding = base64.URLEncoding.WithPadding(base64.NoPadding)

var tempSessionMemoryStore = make(map[string]*webauthn.SessionData)

func (h Handler) BeginRegistration(c *gin.Context) {
	// user := datastore.GetUser() // Find or create the new user
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
	b64UserID := b64Encoding.EncodeToString(user.WaID)
	tempSessionMemoryStore[b64UserID] = session

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
