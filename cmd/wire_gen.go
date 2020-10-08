// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/infinity-oj/cli/internal/app"
	"github.com/infinity-oj/cli/internal/clients"
	"github.com/infinity-oj/cli/internal/clients/accounts"
	"github.com/infinity-oj/cli/internal/clients/submissions"
	"github.com/infinity-oj/cli/internal/clients/volumes"
	accounts2 "github.com/infinity-oj/cli/internal/commands/accounts"
	submissions2 "github.com/infinity-oj/cli/internal/commands/submissions"
	volumes2 "github.com/infinity-oj/cli/internal/commands/volumes"
	"github.com/infinity-oj/cli/internal/config"
	"github.com/infinity-oj/cli/internal/service"
	"github.com/urfave/cli/v2"
)

// Injectors from wire.go:

func CreateApp(cf string) (*cli.App, error) {
	viper, err := config.New(cf)
	if err != nil {
		return nil, err
	}
	options, err := clients.NewOptions(viper)
	if err != nil {
		return nil, err
	}
	client := clients.NewClient(options)
	accountsClient := accounts.NewAccountClient(client)
	accountService := service.NewAccountService(accountsClient)
	accountCommand := accounts2.NewAccountsCommands(accountService)
	volumeClient := volumes.NewVolumeClient(client)
	volumeService := service.NewFileService(volumeClient)
	volumeCommand := volumes2.NewVolumeCommands(volumeService)
	submissionsClient := submissions.NewSubmissionClient(client)
	submissionService := service.NewSubmissionService(submissionsClient)
	submissionCommand := submissions2.NewSubmissionCommands(submissionService)
	cliApp := app.NewApp(accountCommand, volumeCommand, submissionCommand)
	return cliApp, nil
}

// wire.go:

var providerSet = wire.NewSet(config.ProviderSet, accounts2.ProviderSet, volumes2.ProviderSet, submissions2.ProviderSet, service.ProviderSet, app.ProviderSet, clients.ProviderSet)
