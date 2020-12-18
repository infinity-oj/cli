package judgements

import (
	"fmt"
	"github.com/infinity-oj/server-v2/pkg/api"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type JudgementCommands *cli.Command

var ProviderSet = wire.NewSet(NewJudgementsCommands)

func NewJudgementsCommands(api api.API) JudgementCommands {
	var subCommands = []*cli.Command{
		NewCreateJudgementCommand(api),
		NewQueryJudgementCommand(api),
		//NewCreateSubmissionCommand(submissionService),
		//NewDispatchSubmissionCommand(submissionService),
	}
	return &cli.Command{
		Name:        "judgements",
		Aliases:     []string{"j"},
		Usage:       "options for judgements actions",
		UsageText:   "",
		Description: "",
		ArgsUsage:   "",
		Category:    "",
		BashComplete: func(c *cli.Context) {
			// This will complete if no args are passed
			if c.NArg() > 0 {
				return
			}
			for _, t := range subCommands {
				fmt.Println(t.Name)
			}
		},
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		OnUsageError:           nil,
		Subcommands:            subCommands,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
