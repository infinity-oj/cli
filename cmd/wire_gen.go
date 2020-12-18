// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/infinity-oj/cli/internal/app"
	"github.com/infinity-oj/cli/internal/clients"
	"github.com/infinity-oj/cli/internal/commands/accounts"
	"github.com/infinity-oj/cli/internal/commands/judgements"
	"github.com/infinity-oj/cli/internal/commands/submissions"
	"github.com/infinity-oj/cli/internal/commands/volumes"
	"github.com/infinity-oj/cli/internal/config"
	"github.com/urfave/cli/v2"
)

// Injectors from wire.go:

func CreateApp() (*cli.App, error) {
	viper, err := config.New()
	if err != nil {
		return nil, err
	}
	api := clients.NewClient(viper)
	accountCommands := accounts.NewAccountsCommands(api)
	volumeCommands := volumes.NewVolumeCommands(api)
	submissionCommands := submissions.NewSubmissionCommands(api)
	judgementCommands := judgements.NewJudgementsCommands(api)
	cliApp := app.NewApp(accountCommands, volumeCommands, submissionCommands, judgementCommands)
	return cliApp, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, accounts.ProviderSet, volumes.ProviderSet, submissions.ProviderSet, judgements.ProviderSet, app.ProviderSet, clients.ProviderSet)
