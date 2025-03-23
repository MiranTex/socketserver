package Classes

type HttpRequest struct {
	AccessToken     string                 `json:"access_token"`
	EventType       string                 `json:"event_type"`
	EventData       map[string]interface{} `json:"event_data"`
	ClientId        string                 `json:"client_id"`
	ClusterPublicId string                 `json:"cluster_public_id"`
}
