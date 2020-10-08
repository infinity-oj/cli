package volumes

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

func NewCreateVolumeCommand(volumeService service.VolumeService) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new volume",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,

		Action: func(c *cli.Context) error {
			volume, err := volumeService.CreateVolume()
			if err != nil {
				return err
			}
			fmt.Printf("%+v\n", volume)
			return nil
		},

		OnUsageError: nil,
		Subcommands:  nil,
		Flags:        nil,

		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
