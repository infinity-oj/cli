package volumes

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/infinity-oj/server-v2/pkg/models"
)

type VolumeClient interface {
	CreateVolume() (*models.Volume, error)
}

type volume struct {
	client *resty.Client
}

func (a *volume) CreateVolume() (*models.Volume, error) {
	volume := &models.Volume{}

	resp, err := a.client.R().
		SetResult(volume).
		Post("/volume")
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

	return volume, nil
}

func NewVolumeClient(client *resty.Client) VolumeClient {
	return &volume{
		client: client,
	}
}
