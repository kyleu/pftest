// Package queue - Content managed by Project Forge, see [projectforge.md] for details.
package queue

import (
	"github.com/google/uuid"

	"github.com/kyleu/pftest/app/util"
)

type Message struct {
	ID      uuid.UUID `json:"id"`
	Topic   string    `json:"topic"`
	Param   any       `json:"param"`
	Retries int       `json:"retries,omitempty"`
}

func NewMessage(topic string, param any) *Message {
	return &Message{ID: util.UUID(), Topic: topic, Param: param}
}
