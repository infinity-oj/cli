package accounts

import (
	"bufio"
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
	"strings"
	"syscall"

	"github.com/fatih/color"
	"github.com/infinity-oj/cli/internal/services"
	"github.com/rodaine/table"
	"github.com/urfave/cli/v2"
)

func NewCreateAccountCommand(accountService services.AccountService) *cli.Command {
	return &cli.Command{
		Name:         "create",
		Aliases:      []string{"c"},
		Usage:        "create a new account",
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

			fmt.Print("Re-enter password: ")
			bytePassword, err = terminal.ReadPassword(int(syscall.Stdin))
			if err != nil {
				return err
			}
			rePassword := strings.TrimSpace(string(bytePassword))
			fmt.Println()

			if password != rePassword {
				fmt.Println("The two passwords entered are different")
				return nil
			}

			fmt.Print("Enter email: ")
			email, err := reader.ReadString('\n')
			if err != nil {
				return err
			}
			email = strings.TrimSpace(email)

			account, err := accountService.Create(username, password, email)
			if err != nil {
				return err
			}

			headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
			columnFmt := color.New(color.FgYellow).SprintfFunc()

			tbl := table.New("ID", "Time", "Name", "Email")
			tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

			tbl.AddRow(account.ID, account.CreatedAt, account.Name, account.Email)

			tbl.Print()

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
