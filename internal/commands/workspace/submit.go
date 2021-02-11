package workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/infinity-oj/cli/internal/output"

	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewSubmitCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "submit",
		Usage:        "",
		UsageText:    "",
		Description:  "submit current workspace to corresponding problem",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			pwd, err := os.Getwd()
			if err != nil {
				return err
			}
			fmt.Println(pwd)

			iojPath := filepath.Join(pwd, ".ioj")
			if isExist, err := exists(iojPath); err != nil {
				return err
			} else {
				if !isExist {
					return errors.New("not a ioj workspace")
				}
			}

			volumeAPI := api.NewVolumeAPI()

			volume, err := volumeAPI.CreateVolume()
			if err != nil {
				return err
			}

			err = filepath.Walk(pwd, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if strings.HasPrefix(path, filepath.Join(pwd, ".ioj")) {
					return err
				}
				if strings.HasPrefix(path, filepath.Join(pwd, "resources")) {
					return err
				}
				if strings.HasPrefix(path, filepath.Join(pwd, ".gitignore")) {
					return err
				}
				if path == pwd {
					return err
				}
				fmt.Println(path)

				if info.IsDir() {
					volume, err = volumeAPI.CreateDirectory(volume.Name, strings.TrimPrefix(path, pwd))
				} else {
					fileBytes, err := ioutil.ReadFile(path)
					if err != nil {
						return err
					}
					volume, err = volumeAPI.CreateFile(volume.Name, strings.TrimPrefix(path, pwd), fileBytes)
				}
				if err != nil {
					return err
				}

				return err
			})

			byteValue, _ := ioutil.ReadFile(filepath.Join(iojPath, "config.json"))
			config := Config{}
			if err := json.Unmarshal(byteValue, &config); err != nil {
				return err
			}

			code, submission, _, err := api.NewSubmissionAPI().Create(config.ProblemName, volume.Name)
			if err != nil {
				return err
			}

			if code == http.StatusOK {
				tbl := output.NewTable("ID", "Time", "Problem", "Volume")
				tbl.AddRow(submission.Name, submission.CreatedAt, submission.ProblemId, submission.UserVolume)
				tbl.Print()
			} else if code == http.StatusForbidden {
				fmt.Println("Rejected")
			} else {
				fmt.Printf("Server aborts with status code %d\n", code)
			}

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
