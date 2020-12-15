package judgements

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/infinity-oj/cli/internal/services"
	"github.com/urfave/cli/v2"
)

func NewQueryJudgementCommand(judgementService services.JudgementService) *cli.Command {
	return &cli.Command{
		Name:         "query",
		Aliases:      []string{"q"},
		Usage:        "query judgements",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			judgements, err := judgementService.Query()
			if err != nil {
				return err
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Time", "Submission", "Score")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			for _, judgement := range judgements {
				tbl.AddRow(judgement.JudgementId, judgement.CreatedAt, judgement.SubmissionId, judgement.Score)
			}

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
