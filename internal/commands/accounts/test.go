package accounts

import (
	"github.com/infinity-oj/cli/internal/output"
	"github.com/infinity-oj/cli/internal/services"
	"github.com/urfave/cli/v2"
)

func NewTestAccountCommand(accountService services.AccountService) *cli.Command {
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

			tbl := output.NewTable("ID", "Time", "Name", "Email")
			tbl.AddRow(account.ID, account.CreatedAt, account.Name, account.Email)
			tbl.Print()

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
