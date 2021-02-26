package problems

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/output"
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewQueryProblemCommand(api api.API) *cli.Command {
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
			judgementId := c.String("judgement")

			if judgementId == "" {
				judgements, err := api.NewJudgementAPI().QueryJudgements()
				if err != nil {
					return err
				}

				tbl := output.NewTable("ID", "Time", "Submission", "Score")
				for _, judgement := range judgements {
					tbl.AddRow(judgement.JudgementId, judgement.CreatedAt, judgement.SubmissionId, judgement.Score)
				}
				tbl.Print()
			} else {

				judgement, err := api.NewJudgementAPI().QueryJudgement(judgementId)
				if err != nil {
					return err
				}

				tbl := output.NewTable("ID", "Time", "Submission", "Score")
				tbl.AddRow(judgement.JudgementId, judgement.CreatedAt, judgement.SubmissionId, judgement.Score)
				tbl.Print()
				//
				//var msg struct {
				//	Warning string `json:"warning"`
				//	Error   string `json:"error"`
				//}
				//
				//err = json.Unmarshal([]byte(judgement.Msg), &msg)
				//if err != nil {
				//
				//} else {
				//	fmt.Println("Warning:")
				//	fmt.Println(msg.Warning)
				//	fmt.Println()
				//	fmt.Println("Error:")
				//	fmt.Println(msg.Error)
				//}
				fmt.Println()

				fmt.Println(judgement.Msg)
			}

			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "judgement",
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
