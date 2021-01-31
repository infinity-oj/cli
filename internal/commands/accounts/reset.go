package accounts

import (
	"bufio"
	"fmt"
	"github.com/infinity-oj/server-v2/pkg/api"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"

	"github.com/urfave/cli/v2"
)

func NewResetAccountCommand(api api.API) *cli.Command {
	return &cli.Command{
		Name:         "reset",
		Aliases:      []string{"r"},
		Usage:        "reset account password",
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

			fmt.Print("Enter old password: ")
			bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			password := strings.TrimSpace(string(bytePassword))
			fmt.Println()

			fmt.Print("Enter new password: ")
			bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			newPassword := strings.TrimSpace(string(bytePassword))
			fmt.Println()

			fmt.Print("Re-enter new password: ")
			bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			reNewPassword := strings.TrimSpace(string(bytePassword))
			fmt.Println()

			if newPassword != reNewPassword {
				fmt.Println("The two passwords entered are different")
				return nil
			}

			err = api.NewAccountAPI().ResetCredential(username, password, newPassword)
			if err != nil {
				return err
			}

			fmt.Println("success")

			return nil
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
