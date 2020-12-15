// +build wireinject

package main

import (
	"github.com/infinity-oj/cli/internal/app"
	"github.com/infinity-oj/cli/internal/clients"

	"github.com/infinity-oj/cli/internal/commands/accounts"
	"github.com/infinity-oj/cli/internal/commands/judgements"
	"github.com/infinity-oj/cli/internal/commands/submissions"
	"github.com/infinity-oj/cli/internal/commands/volumes"
	"github.com/infinity-oj/cli/internal/config"
	"github.com/infinity-oj/cli/internal/services"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

var providerSet = wire.NewSet(
	config.ProviderSet,
	accounts.ProviderSet,
	volumes.ProviderSet,
	//problem.ProviderSet,
	submissions.ProviderSet,
	judgements.ProviderSet,
	services.ProviderSet,
	app.ProviderSet,

	clients.ProviderSet,
)

func CreateApp() (*cli.App, error) {
	panic(wire.Build(providerSet))
}
