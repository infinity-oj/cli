package accounts

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

func NewCreateAccountCommand(accountService service.AccountService) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new accounts",
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
			email := c.String("email")
			account, err := accountService.Create(username, password, email)
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", account)
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "username",
				Required: true,
				Aliases:  []string{"u", "username"},
				Usage:    "username for new accounts",
			},
			&cli.StringFlag{
				Name:     "email",
				Required: true,
				Aliases:  []string{"e"},
				Usage:    "email for new account",
			},
			&cli.StringFlag{
				Name:     "password",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "password for new account",
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
