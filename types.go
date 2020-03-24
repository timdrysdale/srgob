package srgob

type Message struct {
	Topic        string
	Data         []byte
	ConnectionID string
}
