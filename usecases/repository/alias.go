package repository

import (
	"github.com/Ras96/traq-kinano-cli/ent"
	"github.com/gofrs/uuid"
)

type AliasRepository interface {
	CallAlias(short string) (*ent.Alias, error)
	AddAlias(args *AddAliasArgs) (*ent.Alias, error)
}

type AddAliasArgs struct {
	UserID uuid.UUID
	Short  string
	Long   string
}

func NewAddAliasArgs(userID uuid.UUID, short string, long string) *AddAliasArgs {
	return &AddAliasArgs{
		UserID: userID,
		Short:  short,
		Long:   long,
	}
}
