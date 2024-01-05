package emailhelper

type Helper interface {
	Send(to []string, subject string, body string) error
}
