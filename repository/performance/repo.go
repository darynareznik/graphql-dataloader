package performance

import (
	"context"
	"fmt"
	"graphql-dataloader/graph/model"

	"github.com/rs/zerolog/log"
)

type Repo struct {
	data performanceByAthleteID
}

func NewRepo() *Repo {
	return &Repo{data: performanceStub}
}

func (r *Repo) GetPerformance(ctx context.Context, athleteIDs []string) ([]*model.Performance, []error) {
	log.Info().Msgf("getting performance for athlete ids %v", athleteIDs)

	performance := make([]*model.Performance, 0, len(athleteIDs))
	errs := make([]error, 0, len(athleteIDs))
	for _, id := range athleteIDs {
		i, ok := r.data[id]
		if !ok {
			performance = append(performance, nil)
			errs = append(errs, fmt.Errorf("performance not found for athlete id %v", id))
		} else {
			performance = append(performance, i)
			errs = append(errs, nil)
		}
	}

	return performance, errs
}
