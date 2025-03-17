package Classes

import (
	"encoding/json"
)

type Event struct {
	SenderId  string
	EventType string                 `json:"eventType"`
	Data      map[string]interface{} `json:"data"`
}

func CreateEvent(senderId string, eventType string, data map[string]interface{}) Event {

	return Event{
		SenderId:  senderId,
		EventType: eventType,
		Data:      data,
	}
}

func (event *Event) GetEventType() string {
	return event.EventType
}

func (event *Event) AmItheSender(senderId string) bool {
	return event.SenderId == senderId
}

func (event *Event) ToJson() *string {
	jsonVal, err := json.Marshal(event)
	if err != nil {
		print("Error converting event to json")
		return nil
	}

	jsonString := string(jsonVal)

	return &jsonString
}

func (event *Event) ToByteArray() []byte {
	jsonVal, err := json.Marshal(event)
	if err != nil {
		print("Error converting event to json")
		return nil
	}

	return jsonVal
}
