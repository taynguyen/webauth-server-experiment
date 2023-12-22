package pwl

import (
	"encoding/base64"

	"github.com/go-webauthn/webauthn/webauthn"
)

var b64Encoding = base64.URLEncoding.WithPadding(base64.NoPadding)

var userSessionRepo = make(map[string]*webauthn.SessionData)
var userRepo = make(map[string]WAUser)

func setUserSession(id []byte, value *webauthn.SessionData) {
	idStr := b64Encoding.EncodeToString(id)
	userSessionRepo[idStr] = value
}

func setUser(id []byte, value WAUser) {
	idStr := b64Encoding.EncodeToString(id)
	userRepo[idStr] = value
}
