package athletes

import (
	"context"
	"graphql-dataloader/graph/model"

	"github.com/rs/zerolog/log"
)

type Repo struct {
	data []*model.Athlete
}

func NewRepo() *Repo {
	return &Repo{data: athletesStub}
}

func (repo *Repo) GetAthletes(ctx context.Context) ([]*model.Athlete, error) {
	log.Info().Msg("getting all athletes")
	return repo.data, nil
}
