package model

type Message struct {
	SenderID uint64
	Payload  string
	IsGroup  bool
}
