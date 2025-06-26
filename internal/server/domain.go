package server

import (
	"goapp/internal/domain/health"
	"goapp/internal/domain/profile"
)

func (s *Server) initDomain() {
	s.logger.Info("Application domain init")

	s.healthUsecases = health.NewUsecases(s.db)
	s.profileUsecases = profile.NewUsecases(s.db)
}
