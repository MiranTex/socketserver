package Classes

type HttpClient struct {
	ClientBase
	notifyUrl string
}

func NewHttpClient(id string, notifyUrl string, subscriptions []string) HttpClient {

	return HttpClient{
		ClientBase: ClientBase{
			id:            id,
			Subscriptions: subscriptions,
		},
		notifyUrl: notifyUrl,
	}
}

func (c HttpClient) GetConnectionSatus() bool {
	return true
}

func (c HttpClient) SendEvent(event Event) (bool, error) {
	return true, nil
}
