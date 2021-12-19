/*
Copyright © 2021 Ras96 <asymptote.k.k@gmail.com>

*/
package cmd

import (
	"github.com/Ras96/traq-kinano-cli/util/traq"
	"github.com/gofrs/uuid"
	"github.com/spf13/cobra"
)

func (c *Cmds) aliasCmd() *cobra.Command {
	// aliasCmd represents the alias command
	aliasCmd := &cobra.Command{
		Use:     "alias",
		Short:   "Call aliases",
		Aliases: []string{"a"},
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			alias, err := c.h.CallAlias(c.ctx, args[0])
			if err != nil {
				return err
			}

			traq.MustPostMessage(c.pl.Message.ChannelID, alias.Long)

			return nil
		},
	}

	aliasCmd.AddCommand(c.addAliasCmd())

	return aliasCmd
}

func (c *Cmds) addAliasCmd() *cobra.Command {
	// addAliasCmd represents the addAlias command
	addAliasCmd := &cobra.Command{
		Use:   "add",
		Short: "Add an new alias",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			userID, err := uuid.FromString(c.pl.Message.User.ID)
			if err != nil {
				return err
			}

			if err := c.h.AddAlias(c.ctx, userID, args[0], args[1]); err != nil {
				return err
			}

			traq.MustAddStamp(c.pl.Message.ID, "attoteki_seicho")

			return nil
		},
	}

	return addAliasCmd
}
