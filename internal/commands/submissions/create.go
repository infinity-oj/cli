package submissions

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

func NewCreateSubmissionCommand(submissionService service.SubmissionService) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new submissions",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			problemId := c.String("problemId")
			volume := c.String("volume")

			submission, err := submissionService.Create(problemId, volume)
			if err != nil {
				return err
			}

			fmt.Printf("%+v\n", submission)

			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "problemId",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "problem ID of the submissions",
			},
			&cli.StringFlag{
				Name:     "volume",
				Required: true,
				Aliases:  []string{"v"},
				Usage:    "volume containing submitting code",
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
