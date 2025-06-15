package server

import "goapp/internal/domain/health"

func (s *Server) initDomain() {
	s.logger.Info("REST server domain init")

	s.initHealth()
}

func (s *Server) initHealth() {
	usecase := health.NewUsecase()
	s.health = health.NewHandler(usecase, s.logger)
}
