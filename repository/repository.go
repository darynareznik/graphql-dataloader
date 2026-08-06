package repository

import (
	"context"
	"graphql-dataloader/graph/model"
)

type PerformanceRepo interface {
	GetPerformance(ctx context.Context, athleteIDs []string) ([]*model.Performance, []error)
}

type AthletesRepo interface {
	GetAthletes(ctx context.Context) ([]*model.Athlete, error)
}
