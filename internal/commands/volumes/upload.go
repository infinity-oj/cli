package volumes

import (
	"fmt"
	"github.com/infinity-oj/server-v2/pkg/api"
	"io/ioutil"
	"path"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func uploadFile(api api.API, base, localFilePath, volume, remoteDir string) (err error) {
	dat, err := ioutil.ReadFile(path.Join(base, localFilePath))
	if err != nil {
		return err
	}

	err = api.NewVolumeAPI().CreateFile(volume, path.Join(remoteDir, path.Base(localFilePath)), dat)
	if err != nil {
		return
	}
	return
}

func uploadDirectory(api api.API, base, localDir, volume, remoteDir string) (err error) {
	files, err := ioutil.ReadDir(path.Join(base, localDir))
	if err != nil {
		return
	}
	remoteDir = path.Join(remoteDir, localDir)

	err = api.NewVolumeAPI().CreateDirectory(volume, remoteDir)
	if err != nil {
		return
	}

	for _, f := range files {
		if f.IsDir() {
			err = uploadDirectory(api, base, path.Join(localDir, f.Name()), volume, remoteDir)
		} else {
			err = uploadFile(api, base, path.Join(localDir, f.Name()), volume, remoteDir)
		}
		if err != nil {
			return
		}
	}
	return
}

func NewUploadCommand(api api.API) *cli.Command {
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
			vp := c.String("volumePath")
			p := c.String("path")
			r := c.Bool("recursive")

			p, err := filepath.Abs(p)
			if err != nil {
				return err
			}
			base := filepath.Base(p)
			p = filepath.Dir(p)

			if r {
				if err := uploadDirectory(api, p, base, s, vp); err != nil {
					return err
				}
			} else {
				if err := uploadFile(api, p, base, s, vp); err != nil {
					return err
				}
			}

			fmt.Println("success!")

			return nil
		},

		OnUsageError: nil,
		Subcommands:  nil, Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "volume",
				Required: true,
				Aliases:  []string{"v"},
				Usage:    "target volume(path) you want to upload",
			},
			&cli.StringFlag{
				Name:        "volumePath",
				Aliases:     []string{"vp"},
				Usage:       "target volume path",
				Required:    false,
				Value:       "/",
				DefaultText: "/",
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
