package handler

import (
	"context"

	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/usecases/repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
	"github.com/gofrs/uuid"
)

type Handlers interface {
	// Alias
	CallAlias(ctx context.Context, short string) (*ent.Alias, error)
	AddAlias(ctx context.Context, userID uuid.UUID, short string, long string) error
	// Ping
	Ping() error
}

type handlers struct {
	Repo repository.Repositories
	Srv  service.Services
}

func NewHandlers(repo repository.Repositories, srv service.Services) Handlers {
	return &handlers{
		Repo: repo,
		Srv:  srv,
	}
}
