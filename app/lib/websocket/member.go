package websocket

import (
	"github.com/google/uuid"
	"github.com/samber/lo"

	"github.com/kyleu/pftest/app/util"
)

type UpdateMemberParams struct {
	ID   uuid.UUID `json:"id"`
	Role string    `json:"role"`
}

func (s *Service) GetAllMembers(key string) []*Connection {
	ch, ok := s.channels[key]
	if !ok {
		return nil
	}
	return lo.FilterMap(ch.ConnIDs, func(cID uuid.UUID, _ int) (*Connection, bool) {
		c, ok := s.connections[cID]
		return c, ok
	})
}

func (s *Service) GetOnline(key string) []uuid.UUID {
	ch, ok := s.channels[key]
	if !ok {
		return nil
	}
	online := make([]uuid.UUID, 0)
	lo.ForEach(ch.ConnIDs, func(cID uuid.UUID, _ int) {
		c, ok := s.connections[cID]
		if ok && c != nil && (!lo.Contains(online, c.Profile.ID)) {
			online = append(online, c.Profile.ID)
		}
	})
	return online
}

func (s *Service) sendOnlineUpdate(ch string, connID uuid.UUID, userID uuid.UUID, connected bool, logger util.Logger) error {
	p := OnlineUpdate{UserID: userID, Connected: connected}
	onlineMsg := NewMessage(&userID, ch, "online-update", p)
	return s.WriteChannel(onlineMsg, logger, connID)
}
