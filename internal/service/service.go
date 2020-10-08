package service

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewAccountService,
	NewFileService,
	NewProblemService,
	NewSubmissionService,
	NewJudgementService,
)
