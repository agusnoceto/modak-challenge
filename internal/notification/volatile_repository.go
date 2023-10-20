package notification

import (
	"github.com/agusnoceto/modak-challenge/internal/model"
	"github.com/google/uuid"
	"time"
)

type userMessages map[string][]model.Message

type VolatileRepository struct {
	data map[model.MessageKey]userMessages
}

func NewVolatileRepository() *VolatileRepository {
	return &VolatileRepository{
		data: map[model.MessageKey]userMessages{},
	}
}

func (r *VolatileRepository) GetByEmailSince(key model.MessageKey, email string, since time.Time) ([]model.Message, error) {
	messagesMap, found := r.data[key]
	if !found {
		return nil, nil
	}

	usrMsgs, found := messagesMap[email]
	if !found {
		return nil, nil
	}

	result := []model.Message{}

	for _, msg := range usrMsgs {
		if msg.SentAt.After(since) {
			result = append(result, msg)
		}
	}

	return result, nil
}

func (r *VolatileRepository) Insert(key model.MessageKey, email string, msg string) error {
	messagesMap, found := r.data[key]
	if !found {
		messagesMap = userMessages{}
		r.data[key] = messagesMap
	}

	usrMsgs, found := messagesMap[email]
	if !found {
		usrMsgs = []model.Message{}
		messagesMap[email] = usrMsgs
	}
	messagesMap[email] = append(messagesMap[email], model.Message{
		Id:     uuid.New(),
		Email:  email,
		Msg:    msg,
		SentAt: time.Now(),
		Status: model.MessageSent,
	})
	return nil
}
