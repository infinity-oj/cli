package workspace

import (
	"github.com/infinity-oj/server-v2/pkg/api"

	"github.com/google/wire"
	"github.com/urfave/cli/v2"
)

type Commands []*cli.Command

var ProviderSet = wire.NewSet(NewWorkspaceCommands)

func NewWorkspaceCommands(api api.API) Commands {
	return []*cli.Command{
		NewInitCommand(api),
	}
}
