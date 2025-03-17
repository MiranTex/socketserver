package interfaces

type ClusterClientInterface interface {
	Id() string
	getAccessToken() string
}
