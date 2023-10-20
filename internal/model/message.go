package model

import (
	"github.com/google/uuid"
	"time"
)

const (
	MessagePending   MessageStatus = 1
	MessageSent      MessageStatus = 2
	MessageFail      MessageStatus = 3
	MessageDiscarded MessageStatus = 4

	MessageKeyStatus    MessageKey = "status"
	MessageKeyNews      MessageKey = "news"
	MessageKeyMarketing MessageKey = "marketing"
)

type MessageKey string

type MessageStatus int8

type Message struct {
	Id     uuid.UUID
	Email  string
	Msg    string
	SentAt time.Time
	Status MessageStatus
}
