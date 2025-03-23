package interfaces

import Classes "socket_project/classes"

type ClusterClientInterface interface {
	Id() string
	// getAccessToken() string
	SendEvent(event Classes.Event) (bool, error)
}
