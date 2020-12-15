package submissions

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/services"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type SubmissionCommands *cli.Command

var ProviderSet = wire.NewSet(NewSubmissionCommands)

func NewSubmissionCommands(submissionService services.SubmissionService) SubmissionCommands {
	var subCommands = []*cli.Command{
		NewCreateSubmissionCommand(submissionService),
		//NewDispatchSubmissionCommand(submissionService),
	}
	return &cli.Command{
		Name:        "submissions",
		Aliases:     []string{"s"},
		Usage:       "options for submissions actions",
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
