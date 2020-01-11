package session

import "github.com/NatnaelBerhanu-1/tenahub/TenaHub/api/entity"

// SessionService specifies logged in user session related service
type SessionService interface {
	Session(sessionID string) (*entity.Session, []error)
	StoreSession(session *entity.Session) (*entity.Session, []error)
	DeleteSession(sessionID string) (*entity.Session, []error)
}
