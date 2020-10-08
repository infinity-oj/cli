// +build wireinject

package main

import (
	"github.com/infinity-oj/cli/internal/app"
	"github.com/infinity-oj/cli/internal/clients"

	//"github.com/infinity-oj/cli/internal/commands/file"
	//"github.com/infinity-oj/cli/internal/commands/judgement"
	//"github.com/infinity-oj/cli/internal/commands/problem"
	//"github.com/infinity-oj/cli/internal/commands/submission"
	"github.com/infinity-oj/cli/internal/commands/accounts"
	"github.com/infinity-oj/cli/internal/config"
	"github.com/infinity-oj/cli/internal/service"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	accounts.ProviderSet,
	//file.ProviderSet,
	//problem.ProviderSet,
	//submission.ProviderSet,
	//judgement.ProviderSet,
	service.ProviderSet,
	app.ProviderSet,

	clients.ProviderSet,
)

func CreateApp(cf string) (*cli.App, error) {
	panic(wire.Build(providerSet))
}
