package app

import (
	"fmt"
	"github.com/google/wire"
	"github.com/inconshreveable/go-update"
	"github.com/infinity-oj/cli/internal/commands/accounts"
	"github.com/infinity-oj/cli/internal/commands/judgements"
	"github.com/infinity-oj/cli/internal/commands/submissions"
	"github.com/infinity-oj/cli/internal/commands/volumes"
	"github.com/infinity-oj/cli/internal/commands/workspace"
	"github.com/urfave/cli/v2"
	"net/http"
	"net/url"
	"path"
	"time"
)

func NewApp(
	workspaceCommands workspace.Commands,
	accountsCommand accounts.AccountCommands,
	volumeCommand volumes.VolumeCommands,
//problemCommand problem.AccountCommands,
	submissionCommand submissions.SubmissionCommands,
	judgementCommand judgements.JudgementCommands,
) *cli.App {

	app := &cli.App{
		Name:        "ioj-cli",
		HelpName:    "",
		Usage:       "cli tool to interact with ioj",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "v0.0.2",
		Description: "",
		Commands: append([]*cli.Command{
			accountsCommand,
			volumeCommand,
			//problemCommand.AccountCommands,
			submissionCommand,
			judgementCommand,
			{
				Name:         "upgrade",
				Aliases:      nil,
				Usage:        "",
				UsageText:    "",
				Description:  "",
				ArgsUsage:    "",
				Category:     "",
				BashComplete: nil,
				Before:       nil,
				After:        nil,
				Action: func(context *cli.Context) error {
					u, _ := url.Parse("http://10.20.107.171:2333")
					u.Path = path.Join(u.Path, "assets", "cli", filename())
					resp, err := http.Get(u.String())
					if err != nil {
						return err
					}
					defer resp.Body.Close()
					if err := update.Apply(resp.Body, update.Options{}); err != nil {
						return err
					}
					fmt.Println("success")
					return err
				},
				OnUsageError:           nil,
				Subcommands:            nil,
				Flags:                  nil,
				SkipFlagParsing:        false,
				HideHelp:               false,
				HideHelpCommand:        false,
				Hidden:                 false,
				UseShortOptionHandling: false,
				HelpName:               "",
				CustomHelpTemplate:     "",
			}}, workspaceCommands...),
		Flags:                nil,
		EnableBashCompletion: false,
		HideHelp:             false,
		HideVersion:          false,
		BashComplete:         nil,
		Before:               nil,
		After:                nil,
		Action:               nil,
		CommandNotFound:      nil,
		OnUsageError:         nil,
		Compiled:             time.Now(),
		Authors: []*cli.Author{
			{
				Name:  "Yechang Wu",
				Email: "11711918@mail.sustech.edu.cn",
			},
		},
		Copyright:              "",
		Writer:                 nil,
		ErrWriter:              nil,
		ExitErrHandler:         nil,
		Metadata:               nil,
		ExtraInfo:              nil,
		CustomAppHelpTemplate:  "",
		UseShortOptionHandling: false,
	}

	return app
}

var ProviderSet = wire.NewSet(NewApp)
