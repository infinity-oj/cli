package judgements

import (
	"github.com/infinity-oj/cli/internal/output"
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewCancelJudgementCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "cancel",
		Aliases:      []string{"kill"},
		Usage:        "cancel a judgement",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			judgementId := c.String("judgement id")

			judgement, err := api.NewJudgementAPI().CancelJudgement(judgementId)
			if err != nil {
				return err
			}

			tbl := output.NewTable("ID", "Time", "Submission", "Score")
			tbl.AddRow(judgement.JudgementId, judgement.CreatedAt, judgement.SubmissionId, judgement.Score)
			tbl.Print()

			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "judgement id",
				Aliases: []string{"jid"},
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
