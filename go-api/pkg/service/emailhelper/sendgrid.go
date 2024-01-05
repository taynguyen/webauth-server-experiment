package emailhelper

type implSendgrid struct {
	ApiKey string
}

func NewSendgrid(apiKey string) Helper {
	return &implSendgrid{
		ApiKey: apiKey,
	}
}

func (h implSendgrid) Send(to []string, subject string, body string) error {
	// TODO: implement
	return nil
}
