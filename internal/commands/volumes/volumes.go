package volumes

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type VolumeCommand *cli.Command

var ProviderSet = wire.NewSet(NewVolumeCommands)

func NewVolumeCommands(volumeService service.VolumeService) VolumeCommand {
	var subCommands = []*cli.Command{
		NewCreateVolumeCommand(volumeService),
		NewUploadCommand(volumeService),
	}
	return &cli.Command{
		Name:        "volume",
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
