package service

import (
	"fmt"

	"github.com/infinity-oj/cli/internal/clients/volumes"
	"github.com/infinity-oj/server-v2/pkg/models"

	"github.com/pkg/errors"
)

type VolumeService interface {
	CreateVolume() (*models.Volume, error)
	CreateDirectory(fileSpace, directory string) error
	CreateFile(fileSpace, filename string, file []byte) error
}

type volumeService struct {
	volumeClient volumes.VolumeClient
}

func (d *volumeService) CreateDirectory(volume, directory string) error {

	fmt.Println(volume, directory)
	err := d.volumeClient.CreateDirectory(volume, directory)
	if err != nil {
		return errors.Wrap(err, "create directory error")
	}

	return nil
}

func (d *volumeService) CreateFile(volume, filename string, file []byte) error {

	err := d.volumeClient.CreateFile(volume, filename, file)
	if err != nil {
		return errors.Wrap(err, "Create volume error")
	}
	return nil
}

func (d *volumeService) CreateVolume() (*models.Volume, error) {
	volume, err := d.volumeClient.CreateVolume()
	if err != nil {
		return nil, errors.Wrap(err, "create volume error")
	}
	return volume, nil
}

func NewVolumeService(volumeClient volumes.VolumeClient) VolumeService {
	return &volumeService{
		volumeClient: volumeClient,
	}
}
