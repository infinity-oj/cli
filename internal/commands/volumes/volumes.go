package volumes

import (
	"fmt"
	"github.com/infinity-oj/server-v2/pkg/api"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type VolumeCommands *cli.Command

var ProviderSet = wire.NewSet(NewVolumeCommands)

func NewVolumeCommands(api api.API) VolumeCommands {
	var subCommands = []*cli.Command{
		NewCreateVolumeCommand(api),
		NewUploadCommand(api),
	}
	return &cli.Command{
		Name:        "volumes",
		Aliases:     []string{"v"},
		Usage:       "options for volume actions",
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
