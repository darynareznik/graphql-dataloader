package dataloader

import (
	"context"
	"graphql-dataloader/graph/model"
	"graphql-dataloader/repository"
	"net/http"
	"time"

	"github.com/vikstrous/dataloadgen"
)

type ctxKey string

const (
	key = ctxKey("dataloaders")
)

type performanceReader struct {
	performanceRepo repository.PerformanceRepo
}

// getPerformance implements a batch function that can retrieve many performance by athleteID,
// for use in a dataloader
func (r *performanceReader) getPerformance(ctx context.Context, athleteIDs []string) ([]*model.Performance, []error) {
	return r.performanceRepo.GetPerformance(ctx, athleteIDs)
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	PerformanceLoader *dataloadgen.Loader[string, *model.Performance]
}

// NewLoaders instantiates data loaders for the middleware
func NewLoaders(repo repository.PerformanceRepo) *Loaders {
	// define the data loader
	r := &performanceReader{performanceRepo: repo}
	return &Loaders{
		PerformanceLoader: dataloadgen.NewLoader(r.getPerformance, dataloadgen.WithWait(time.Millisecond)),
	}
}

// Middleware injects data loaders into the context
func Middleware(repo repository.PerformanceRepo, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		loader := NewLoaders(repo)
		r = r.WithContext(context.WithValue(r.Context(), key, loader))
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(key).(*Loaders)
}

func GetPerformanceForOne(ctx context.Context, athleteID string) (*model.Performance, error) {
	loaders := For(ctx)
	return loaders.PerformanceLoader.Load(ctx, athleteID)
}
