package accounts

import (
	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

func NewLoginAccountCommand(accountService service.AccountService) *cli.Command {
	return &cli.Command{
		Name:         "login",
		Aliases:      nil,
		Usage:        "login",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			//fmt.Println("new task template: ", c.Args().First())
			username := c.String("username")
			password := c.String("password")
			err := accountService.Login(username, password)
			if err != nil {
				return err
			}
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Required: true,
				Aliases:  []string{"u"},
				Usage:    "account username to login",
			},
			&cli.StringFlag{
				Name:     "password",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "password for this account",
			},
		},
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
