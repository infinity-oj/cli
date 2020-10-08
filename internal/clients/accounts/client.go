package accounts

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/infinity-oj/server-v2/pkg/models"
)

type Client interface {
	CreateAccount(username, password, email string) (*models.Account, error)
}

type account struct {
	client *resty.Client
}

func (a *account) CreateAccount(username, password, email string) (*models.Account, error) {
	account := &models.Account{}

	request := map[string]interface{}{
		"username": username,
		"password": password,
		"email":    email,
	}

	resp, err := a.client.R().
		SetBody(request).
		SetResult(account).
		Post("/accounts/application")
	if err != nil {
		return nil, err
	}

	// Explore response object
	fmt.Println("Response Info:")
	fmt.Println("  ", resp.Request.URL)
	fmt.Println("  Error      :", err)
	fmt.Println("  Status Code:", resp.StatusCode())
	fmt.Println("  Status     :", resp.Status())
	fmt.Println("  Proto      :", resp.Proto())
	fmt.Println("  Time       :", resp.Time())
	fmt.Println("  Received At:", resp.ReceivedAt())
	fmt.Println("  Body       :\n", resp)
	fmt.Println()

	return account, nil
}

func NewAccountClient(client *resty.Client) Client {
	return &account{
		client: client,
	}
}
