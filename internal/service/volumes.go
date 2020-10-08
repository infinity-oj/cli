package service

import (
	"github.com/infinity-oj/cli/internal/clients/volumes"
	"github.com/infinity-oj/server-v2/pkg/models"

	"github.com/pkg/errors"
)

type VolumeService interface {
	CreateVolume() (*models.Volume, error)
	CreateDirectory(fileSpace, directory string) error
	CreateFile(fileSpace, fileName string, data []byte) error
}

type volumeService struct {
	volumeClient volumes.VolumeClient
}

func (d *volumeService) CreateDirectory(fileSpace, directory string) error {

	//fs, err := d.volumeClient.CreateVolume()
	//if err != nil {
	//	return errors.Wrap(err, "create directory error")
	//}
	//fmt.Println(fs.Status)
	return nil
}

func (d *volumeService) CreateFile(fileSpace, fileName string, data []byte) error {
	//req := &proto.CreateFileRequest{
	//	FileSpace: fileSpace,
	//	FilePath:  fileName,
	//	Data:      data,
	//}
	//
	//fs, err := d.volumeClient.CreateFile(context.TODO(), req)
	//if err != nil {
	//	return errors.Wrap(err, "Create volume error")
	//}
	//fmt.Println(fs.Status)
	return nil
}

func (d *volumeService) CreateVolume() (*models.Volume, error) {
	volume, err := d.volumeClient.CreateVolume()
	if err != nil {
		return nil, errors.Wrap(err, "create volume error")
	}
	return volume, nil
}

func NewFileService(volumeClient volumes.VolumeClient) VolumeService {
	return &volumeService{
		volumeClient: volumeClient,
	}
}
