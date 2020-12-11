package service

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/clients"
	"github.com/infinity-oj/cli/internal/clients/accounts"
	"github.com/infinity-oj/server-v2/pkg/models"
	"github.com/pkg/errors"
)

type AccountService interface {
	Create(username, password, email string) (*models.Account, error)
	Login(username, password string) error
	Test() (*models.Account, error)
}

type accountService struct {
	accountClient accounts.Client
}

func NewAccountService(accountClient accounts.Client) AccountService {
	return &accountService{
		accountClient: accountClient,
	}
}

func (s *accountService) Create(username, password, email string) (*models.Account, error) {

	account, err := s.accountClient.CreateAccount(username, password, email)
	if err != nil {
		return nil, errors.Wrap(err, "create accounts error")
	}

	return account, nil
}

func (s *accountService) Login(username, password string) error {

	err := s.accountClient.Login(username, password)
	if err != nil {
		return errors.Wrap(err, "create accounts error")
	}

	err = clients.Jar.Save()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (s *accountService) Test() (*models.Account, error) {

	account, err := s.accountClient.TestAccount()
	if err != nil {
		return nil, errors.Wrap(err, "create accounts error")
	}
	return account, err
}
