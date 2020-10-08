package submissions

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/infinity-oj/server-v2/pkg/models"
)

type Client interface {
	CreateSubmission(problemId, volume string) (*models.Submission, error)
}

type submission struct {
	client *resty.Client
}

func (a *submission) CreateSubmission(problemId, volume string) (*models.Submission, error) {
	submission := &models.Submission{}

	request := map[string]interface{}{
		"problemId": problemId,
		"volume":    volume,
	}

	resp, err := a.client.R().
		SetBody(request).
		SetResult(submission).
		Post("/submissions")
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

	return submission, nil
}

func NewSubmissionClient(client *resty.Client) Client {
	return &submission{
		client: client,
	}
}
