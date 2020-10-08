package service

import (
	"github.com/infinity-oj/cli/internal/clients/accounts"

	"github.com/infinity-oj/server-v2/pkg/models"
	"github.com/pkg/errors"
)

type AccountService interface {
	Create(username, password, email string) (*models.Account, error)
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
