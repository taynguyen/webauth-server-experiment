package pwl

import "github.com/go-webauthn/webauthn/webauthn"

// TODO: move?
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
