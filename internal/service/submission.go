package service

import (
	"github.com/infinity-oj/cli/internal/clients/submissions"
	"github.com/infinity-oj/server-v2/pkg/models"

	"github.com/pkg/errors"
)

type SubmissionService interface {
	Create(problemId, volume string) (*models.Submission, error)
}

type service struct {
	client submissions.Client
}

func (s *service) Create(problemId, volume string) (*models.Submission, error) {

	submission, err := s.client.CreateSubmission(problemId, volume)
	if err != nil {
		return nil, errors.Wrap(err, "create submissions error")
	}

	return submission, nil
}

func NewSubmissionService(client submissions.Client) SubmissionService {
	return &service{
		client: client,
	}
}
