package submissions

import (
	"fmt"
	"net/http"

	"github.com/infinity-oj/cli/internal/output"
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewCreateSubmissionCommand(api api.API) *cli.Command {
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

			code, submission, _, err := api.NewSubmissionAPI().Create(problemId, volume)
			if err != nil {
				return err
			}

			if code == http.StatusOK {
				tbl := output.NewTable("ID", "Time", "Problem", "Volume")
				tbl.AddRow(submission.Name, submission.CreatedAt, submission.ProblemId, submission.UserVolume)
				tbl.Print()
			} else if code == http.StatusForbidden {
				fmt.Println("Rejected")
			} else {
				fmt.Printf("Server aborts with status code %d\n", code)
			}

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
