package accounts

import (
	"bufio"
	"fmt"
	"github.com/infinity-oj/cli/internal/clients"
	"github.com/infinity-oj/server-v2/pkg/api"
	"github.com/urfave/cli/v2"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"
)

func NewLoginAccountCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "login",
		Aliases:      nil,
		Usage:        "login",
		UsageText:    "",
		Description:  "",
		ArgsUsage:    "",
		Category:     "",
		BashComplete: nil,
		Before:       nil,
		After:        nil,
		Action: func(c *cli.Context) error {
			reader := bufio.NewReader(os.Stdin)

			fmt.Print("Enter username: ")
			username, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			username = strings.TrimSpace(username)

			fmt.Print("Enter password: ")
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			password := strings.TrimSpace(string(bytePassword))
			fmt.Println()

			err = api.NewAccountAPI().Login(username, password)
			if err != nil {
				return err
			}

			fmt.Println("Login successfully")

			return clients.Jar.Save()
		},
		OnUsageError:           nil,
		Subcommands:            nil,
		Flags:                  nil,
		SkipFlagParsing:        false,
		HideHelp:               false,
		Hidden:                 false,
		UseShortOptionHandling: false,
		HelpName:               "",
		CustomHelpTemplate:     "",
	}
}
