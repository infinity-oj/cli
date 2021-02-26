package problems

import (
	"fmt"

	"github.com/infinity-oj/server-v2/pkg/api"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type ProblemCommands *cli.Command

var ProviderSet = wire.NewSet(NewProblemsCommands)

func NewProblemsCommands(api api.API) ProblemCommands {
	var subCommands = []*cli.Command{
		NewCreateProblemCommand(api),
	}
	return &cli.Command{
		Name:        "problem",
		Aliases:     []string{"p"},
		Usage:       "options for problem actions",
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
