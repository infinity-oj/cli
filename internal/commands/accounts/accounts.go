package accounts

import (
	"fmt"
	"github.com/infinity-oj/server-v2/pkg/api"

	"github.com/google/wire"

	"github.com/urfave/cli/v2"
)

type AccountCommands *cli.Command

var ProviderSet = wire.NewSet(NewAccountsCommands)

func NewAccountsCommands(api api.API) AccountCommands {
	var subCommands = []*cli.Command{
		NewCreateAccountCommand(api),
		NewLoginAccountCommand(api),
		NewTestAccountCommand(api),
		NewResetAccountCommand(api),
	}
	return &cli.Command{
		Name:        "accounts",
		Aliases:     []string{"a"},
		Usage:       "options for accounts actions",
		UsageText:   "",
		Description: "",
		ArgsUsage:   "",
		Category:    "",
		BashComplete: func(c *cli.Context) {
			// This will complete if no args are passed
			if c.NArg() > 0 {
				return
			}
			for _, t := range subCommands {
				fmt.Println(t.Name)
			}
		},
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		OnUsageError:           nil,
		Subcommands:            subCommands,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
