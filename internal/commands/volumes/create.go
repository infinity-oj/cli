package volumes

import (
	"github.com/infinity-oj/cli/internal/output"

	"github.com/infinity-oj/cli/internal/services"
	"github.com/urfave/cli/v2"
)

func NewCreateVolumeCommand(volumeService services.VolumeService) *cli.Command {
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

			tbl := output.NewTable("ID", "Time")
			tbl.AddRow(volume.Name, volume.CreatedAt)
			tbl.Print()

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
