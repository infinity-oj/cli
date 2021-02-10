package workspace

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/infinity-oj/cli/internal/output"

	"github.com/infinity-oj/server-v2/pkg/models"

	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
)

func NewInitCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "init",
		Usage:        "<problem name> [<target directory>]",
		UsageText:    "",
		Description:  "initialize a workspace",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			problemName := ""
			if c.NArg() > 0 {
				problemName = c.Args().Get(0)
			} else {
				return errors.New("missing argument <problem name>")
			}

			problem, err := api.NewProblemAPI().GetProblem(problemName)
			if err != nil {
				return err
			}

			pwd, err := os.Getwd()
			if err != nil {
				return err
			}
			path := filepath.Join(pwd, problemName)
			dirname := problemName
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
			if isExist {
				if isEmpty, err = IsDirEmpty(path); err != nil {
					return err
				}
				if !isEmpty {
					return errors.New(fmt.Sprintf(
						"destination path '%s' already exists and is not an empty directory",
						dirname))
				}
			}
			fmt.Println(path)

			if err = InitWorkSpace(problem, path); err != nil {
				return err
			}

			tbl := output.NewTable("ID", "Name", "Title")
			tbl.AddRow(problem.ID, problem.Name, problem.Title)
			tbl.Print()

			fmt.Println("success!")
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

func InitWorkSpace(problem *models.Problem, path string) error {
	mkdir := func(p string) error {
		if err := os.Mkdir(p, 0644); err != nil {
			return err
		}
		return nil
	}

	if err := mkdir(path); err != nil {
		return err
	}
	if err := mkdir(filepath.Join(path, ".ioj")); err != nil {
		return err
	}

	content := []byte(".ioj\n")
	err := ioutil.WriteFile(filepath.Join(path, ".gitignore"), content, 0644)
	if err != nil {
		return err
	}

	config := Config{ProblemName: problem.Name}

	jsonBytes, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(path, ".ioj", "config.json"), jsonBytes, 0644)

	return nil
}
