package submissions

import (
	"github.com/fatih/color"
	"github.com/rodaine/table"

	"github.com/infinity-oj/cli/internal/services"
	"github.com/urfave/cli/v2"
)

func NewCreateSubmissionCommand(submissionService services.SubmissionService) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new submission",
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

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Time", "Problem", "Volume")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			tbl.AddRow(submission.SubmissionId, submission.CreatedAt, submission.ProblemId, submission.UserVolume)

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
