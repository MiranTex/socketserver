package interfaces

import (
	Classes "socket_project/classes"
)

type ClientInterface interface {
	Id() string
	GetConnectionSatus() bool
	SendEvent(event Classes.Event) (bool, error)
	String() string
}
