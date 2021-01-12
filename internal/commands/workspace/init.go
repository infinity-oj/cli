package workspace

import (
	"errors"
	"fmt"
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"path/filepath"
)

func NewInitCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "init",
		Usage:        "<problem> [<dir>]",
		UsageText:    "",
		Description:  "initialize a workspace",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			problem := ""
			if c.NArg() > 0 {
				problem = c.Args().Get(0)
			}

			pwd, err := os.Getwd()
			if err != nil {
				return err
			}
			path := filepath.Join(pwd, problem)
			dirname := problem
			if c.NArg() > 1 {
				dirname = c.Args().Get(1)
				if filepath.IsAbs(dirname) {
					path = dirname
				} else {
					path = filepath.Join(pwd, dirname)
				}
			}

			isExist := false
			isEmpty := false
			if isExist, err = exists(path); err != nil {
				return err
			}
			if isEmpty, err = IsDirEmpty(path); err != nil {
				return err
			}
			if isExist && !isEmpty {
				return errors.New(fmt.Sprintf(
					"destination path '%s' already exists and is not an empty directory",
					dirname))
			}

			fmt.Println(path)

			return nil
		},
		OnUsageError: nil,
		Subcommands:  nil,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Required:    false,
				Hidden:      false,
				Value:       false,
				DefaultText: "",
				Destination: nil,
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

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IsDirEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, err
	}
	defer f.Close()

	// read in ONLY one file
	_, err = f.Readdir(1)

	// and if the file is EOF... well, the dir is empty.
	if err == io.EOF {
		return true, nil
	}
	return false, err
}
