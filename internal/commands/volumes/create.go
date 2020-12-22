package volumes

import (
	"archive/zip"
	"github.com/infinity-oj/cli/internal/output"
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
	"io"
	"io/ioutil"
)

func NewCreateVolumeCommand(api api.API) *cli.Command {
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
			api.NewAccountAPI().Login("wycers", "Ccat2683")

			zipPath := c.String("zip")

			volume, err := api.NewVolumeAPI().CreateVolume()
			if err != nil {
				return err
			}

			if zipPath != "" {
				r, err := zip.OpenReader(zipPath)
				if err != nil {
					return err
				}
				defer r.Close()

				for _, f := range r.File {
					if !f.FileInfo().IsDir() {
						continue
					}
					err = api.NewVolumeAPI().CreateDirectory(volume.Name, f.Name)
					if err != nil {
						return err
					}
				}

				for _, f := range r.File {
					if f.FileInfo().IsDir() {
						continue
					}
					rc, err := f.Open()
					if err != nil {
						return err
					}
					data, err := ioutil.ReadAll(rc)
					if len(data) == 0 {
						data = []byte{0x0}
					}
					if err != nil {
						return err
					}
					err = api.NewVolumeAPI().CreateFile(volume.Name, f.Name, data)
					if err != nil && err != io.EOF {
						return err
					}
				}
			}

			tbl := output.NewTable("ID")
			tbl.AddRow(volume.Name)
			tbl.Print()

			return nil

		},

		OnUsageError: nil,
		Subcommands:  nil,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "zip",
				Usage: "Create a volume by contents in a zip",
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
