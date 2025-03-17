package Classes

type SocketRequest struct {
	Id            string                 `json:"id"`
	RequestType   string                 `json:"type"`
	EventType     string                 `json:"eventType"`
	EventMessage  map[string]interface{} `json:"eventMessage"`
	Subscriptions []string               `json:"subscriptions"`
}
