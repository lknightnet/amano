package model

type Message struct {
	ID               int                    `json:"id"`
	Level            string                 `json:"level"`
	Text             string                 `json:"text"`
	MicroServiceName string                 `json:"ms_name"`
	Time             string                 `json:"time"`
	Other            map[string]interface{} `json:"other,omitempty"`
}
