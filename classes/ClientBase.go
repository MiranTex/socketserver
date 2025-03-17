package Classes

import "fmt"

type ClientBase struct {
	id            string
	Subscriptions []string `json:"subscriptions"`
}

func (c ClientBase) String() string {
	return fmt.Sprintf("{id: %s, Subscriptions: %v}", c.id, c.Subscriptions)
}

func (c ClientBase) Id() string {
	return c.id
}

func (c Client) IsIdEqualTo(id string) bool {
	return c.id == id
}
