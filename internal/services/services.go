package services

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAccountService,
	NewVolumeService,
	NewSubmissionService,
	NewJudgementService,
)
