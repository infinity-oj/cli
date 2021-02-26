package problems

import (
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewCreateProblemCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new problem with a default page",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,

		Action: func(c *cli.Context) error {
			////fmt.Println("new task template: ", c.Args().First())
			//title := c.String("title")
			//locale := c.String("locale")
			//_, err := api.NewProblemAPI().CreateProblem(title, locale)
			//return err
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "title",
				Required: true,
				Aliases:  []string{"t"},
				Usage:    "title for this problem",
			},
			&cli.StringFlag{
				Name:     "locale",
				Required: true,
				Aliases:  []string{"l"},
				Usage:    "locale of this problem",
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
