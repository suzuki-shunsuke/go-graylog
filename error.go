package graylog

type Error struct {
	Type    string `json:"type"`
	Message string `json:"message"`
}
