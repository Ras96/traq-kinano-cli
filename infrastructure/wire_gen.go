// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package infrastructure

import (
	"context"
	"github.com/Ras96/traq-kinano-cli/cmd"
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/Ras96/traq-kinano-cli/interfaces/handler"
	"github.com/Ras96/traq-kinano-cli/interfaces/repository"
	"github.com/Ras96/traq-kinano-cli/usecases/service"
	"github.com/traPtitech/traq-bot"
)

import (
	_ "github.com/go-sql-driver/mysql"
)

// Injectors from wire.go:

func injectCmds(ctx context.Context, client *ent.Client, pl *traqbot.MessageCreatedPayload) *cmd.Cmds {
	repositories := repository.NewRepositories(client)
	services := service.NewServices()
	handlers := handler.NewHandlers(repositories, services)
	cmds := cmd.NewCmds(ctx, handlers, pl)
	return cmds
}
