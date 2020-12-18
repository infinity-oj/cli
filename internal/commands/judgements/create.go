package judgements

import (
	"github.com/infinity-oj/cli/internal/output"
	"github.com/infinity-oj/cli/internal/services"
	"github.com/urfave/cli/v2"
)

func NewCreateJudgementCommand(judgementService services.JudgementService) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new judgement",
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

			judgement, err := judgementService.Create(problemId, volume)
			if err != nil {
				return err
			}

			tbl := output.NewTable("ID", "Time", "Submission", "Score")
			tbl.AddRow(judgement.JudgementId, judgement.CreatedAt, judgement.SubmissionId, judgement.Score)
			tbl.Print()

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
