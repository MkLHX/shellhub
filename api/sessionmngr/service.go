package sessionmngr

import (
	"context"
	"strings"

	"github.com/shellhub-io/shellhub/api/store"
	"github.com/shellhub-io/shellhub/pkg/api/paginator"
	"github.com/shellhub-io/shellhub/pkg/models"
)

type Service interface {
	ListSessions(ctx context.Context, pagination paginator.Query) ([]models.Session, int, error)
	GetSession(ctx context.Context, uid models.UID) (*models.Session, error)
	CreateSession(ctx context.Context, session models.Session) (*models.Session, error)
	DeactivateSession(ctx context.Context, uid models.UID) error
	SetSessionAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error
	RecordSession(ctx context.Context, uid models.UID, recordString string) error
}

type service struct {
	store store.Store
}

func NewService(store store.Store) Service {
	return &service{store}
}

func (s *service) ListSessions(ctx context.Context, pagination paginator.Query) ([]models.Session, int, error) {
	return s.store.ListSessions(ctx, pagination)
}

func (s *service) GetSession(ctx context.Context, uid models.UID) (*models.Session, error) {
	return s.store.GetSession(ctx, uid)
}

func (s *service) CreateSession(ctx context.Context, session models.Session) (*models.Session, error) {
	return s.store.CreateSession(ctx, session)
}

func (s *service) DeactivateSession(ctx context.Context, uid models.UID) error {
	return s.store.DeactivateSession(ctx, uid)
}

func (s *service) SetSessionAuthenticated(ctx context.Context, uid models.UID, authenticated bool) error {
	return s.store.SetSessionAuthenticated(ctx, uid, authenticated)
}
func (s *service) RecordSession(ctx context.Context, uid models.UID, recordString string) error {
	result := strings.Split(strings.Replace(recordString, "\r\n", "\n", -1), "\n")
	for i := range result {
		if err := s.store.RecordSession(ctx, uid, result[i]); err != nil {
			return err
		}

	}
	return nil

}