package app

import (
	"time"

	"github.com/google/wire"
	"github.com/infinity-oj/cli/internal/commands/accounts"
	"github.com/infinity-oj/cli/internal/commands/judgements"
	"github.com/infinity-oj/cli/internal/commands/submissions"
	"github.com/infinity-oj/cli/internal/commands/volumes"
	"github.com/urfave/cli/v2"
)

func NewApp(
	accountsCommand accounts.AccountCommands,
	volumeCommand volumes.VolumeCommands,
//problemCommand problem.AccountCommands,
	submissionCommand submissions.SubmissionCommands,
	judgementCommand judgements.JudgementCommands,
) *cli.App {
	app := &cli.App{
		Name:        "",
		HelpName:    "",
		Usage:       "",
		UsageText:   "",
		ArgsUsage:   "",
		Version:     "",
		Description: "",
		Commands: []*cli.Command{
			accountsCommand,
			volumeCommand,
			//problemCommand.AccountCommands,
			submissionCommand,
			judgementCommand,
		},
		Flags:                  nil,
		EnableBashCompletion:   false,
		HideHelp:               false,
		HideVersion:            false,
		BashComplete:           nil,
		Before:                 nil,
		After:                  nil,
		Action:                 nil,
		CommandNotFound:        nil,
		OnUsageError:           nil,
		Compiled:               time.Time{},
		Authors:                nil,
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
