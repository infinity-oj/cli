package accounts

import (
	"fmt"
	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

func NewTestAccountCommand(accountService service.AccountService) *cli.Command {
	return &cli.Command{
		Name:         "test",
		Aliases:      nil,
		Usage:        "test",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			account, err := accountService.Test()
			if err != nil {
				return err
			}
			fmt.Println(account)
			return nil
		},
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
