package resolver

import (
	"graphql-dataloader/repository"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require
// here.

type Resolver struct {
	performanceRepo repository.PerformanceRepo
	athletesRepo    repository.AthletesRepo
}

func NewResolver(performanceRepo repository.PerformanceRepo, athletesRepo repository.AthletesRepo) *Resolver {
	return &Resolver{
		performanceRepo: performanceRepo,
		athletesRepo:    athletesRepo,
	}
}
