package server

import (
	"goapp/internal/domain/health"
	"goapp/internal/domain/profile"
	"goapp/internal/domain/profile2"
	profile2repos "goapp/internal/domain/profile2/repositories"
)

func (s *Server) initDomain() {
	s.logger.Info("Application domain init")

	// Грязная версия profile
	s.healthUsecases = health.NewUsecases(s.db)
	s.profileUsecases = profile.NewUsecases(s.db)

	// Чистая версия profile
	profile2Repo := profile2repos.NewPostgresRepository(s.db)
	s.profile2Usecases = profile2.NewUsecases(profile2Repo)
}
