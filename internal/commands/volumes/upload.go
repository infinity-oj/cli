package volumes

import (
	"io/ioutil"
	"path"

	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

func uploadFile(volumeService service.VolumeService, localFilePath, volume, remoteDir string) (err error) {
	dat, err := ioutil.ReadFile(localFilePath)
	if err != nil {
		return err
	}

	err = volumeService.CreateFile(volume, path.Join(remoteDir, path.Base(localFilePath)), dat)
	if err != nil {
		return
	}
	return
}

func uploadDirectory(volumeService service.VolumeService, base, localDir, volume, remoteDir string) (err error) {
	files, err := ioutil.ReadDir(path.Join(base, localDir))
	if err != nil {
		return
	}

	err = volumeService.CreateDirectory(volume, localDir)
	if err != nil {
		return
	}

	for _, f := range files {
		if f.IsDir() {
			if err = uploadDirectory(volumeService, base, path.Join(localDir, f.Name()), volume, path.Join(remoteDir, f.Name())); err != nil {
				return
			}
		} else {
			if err = uploadFile(volumeService, path.Join(base, localDir, f.Name()), volume, remoteDir); err != nil {
				return
			}
		}
	}
	return
}

func NewUploadCommand(fileService service.VolumeService) *cli.Command {
	return &cli.Command{
		Name:         "upload",
		Aliases:      []string{"up"},
		Usage:        "upload a local file or a local directory to a remote volume",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,

		Action: func(c *cli.Context) error {
			//fmt.Println("new task template: ", c.Args().First())
			s := c.String("volume")
			p := c.String("path")
			r := c.Bool("recursive")
			if r {
				if err := uploadDirectory(fileService, p, "", s, ""); err != nil {
					return err
				}
			} else {
				if err := uploadFile(fileService, p, s, ""); err != nil {
					return err
				}
			}

			return nil
		},

		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "volume",
				Required: true,
				Aliases:  []string{"v"},
				Usage:    "target volume you want to upload",
			},
			&cli.StringFlag{
				Name:     "path",
				Required: true,
				Aliases:  []string{"p"},
				Usage:    "volume or directory you want to upload",
			},
			&cli.BoolFlag{
				Name:     "recursive",
				Required: false,
				Aliases:  []string{"r", "R"},
				Usage:    "upload directories and their contents recursively",
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
