package submissions

import (
	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

type DispatchSubmissionCommand = cli.Command

func NewDispatchSubmissionCommand(submissionService service.SubmissionService) *DispatchSubmissionCommand {
	return &DispatchSubmissionCommand{
		Name:         "dispatch",
		Aliases:      []string{"d"},
		Usage:        "dispatch judgement of a submissions",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			//submissionId := c.String("submissions ID")

			//judgementId, err := submissionService.DispatchJudgement(submissionId)
			//if err != nil {
			//	return err
			//}
			//fmt.Println(judgementId)
			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "submissions ID",
				Required: true,
				Aliases:  []string{"s", "sid"},
				Usage:    "submissions ID",
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
