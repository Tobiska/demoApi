package session

import "demoApi/internal/app/domain/group"

type Session struct {
	UUID  string `json:"uuid,omitempty"`
	Group *group.Group
}
