package notification

import (
	"github.com/agusnoceto/modak-challenge/internal/model"
	"time"
)

type Repository interface {
	GetByEmailSince(key model.MessageKey, email string, since time.Time) ([]model.Message, error)
	Insert(key model.MessageKey, email string, msg string) error
}

type Service interface {
	Send(key model.MessageKey, email string, msgs string) error
}

type Limiter interface {
	Allow(key model.MessageKey, email string) (bool, error)
}
