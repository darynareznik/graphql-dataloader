package server

import (
	"context"
	"errors"
	"fmt"
	"graphql-dataloader/config"
	"graphql-dataloader/graph/dataloader"
	"graphql-dataloader/graph/generated"
	"graphql-dataloader/graph/resolver"
	"graphql-dataloader/repository"
	"graphql-dataloader/repository/athletes"
	"graphql-dataloader/repository/performance"
	"net/http"

	gqlhandler "github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/rs/zerolog/log"
)

type Server struct {
	http             *http.Server
	resolver         generated.ResolverRoot
	runErr           error
	readiness        bool
	enablePlayground bool
}

func CreateAndRun(ctx context.Context, cfg config.Config) *Server {
	performanceRepo := performance.NewRepo()
	resolver := resolver.NewResolver(performanceRepo, athletes.NewRepo())

	svr := &Server{
		http: &http.Server{
			Addr: fmt.Sprintf(":%d", cfg.GraphQLPort),
		},
		resolver:         resolver,
		enablePlayground: cfg.GraphQLEnablePlayground,
	}

	svr.setupHandlers(performanceRepo)
	svr.run()

	return svr
}

func (s *Server) setupHandlers(performanceRepo repository.PerformanceRepo) {
	svr := gqlhandler.New(
		generated.NewExecutableSchema(
			generated.Config{
				Resolvers: s.resolver,
			},
		),
	)
	svr.AddTransport(transport.POST{})

	handler := http.NewServeMux()
	handler.Handle("/query", dataloader.Middleware(performanceRepo, svr))

	if s.enablePlayground {
		handler.Handle("/", playground.Handler("GraphQL Playground", "/query"))
	}

	s.http.Handler = handler
}

func (s *Server) run() {
	log.Info().Msg("graphql service: starting")

	go func() {
		log.Debug().Msgf("graphql service: addr=%v", s.http.Addr)
		if err := s.http.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			s.runErr = err
			log.Err(err).Msg("graphql service: error starting")
		}
	}()

	s.readiness = true
}

func (s *Server) Close(ctx context.Context) {
	if err := s.http.Shutdown(ctx); err != nil {
		log.Err(err).Msg("graphql service shutdown")
	}

	log.Info().Msg("graphql service: stopped")
}
