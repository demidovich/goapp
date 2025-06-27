package server

import (
	"goapp/internal/domain/profile"
	"goapp/internal/domain/profile2"
	"goapp/internal/utility/rest/request"
	"goapp/internal/utility/rest/respond"
	"net/http"
)

func (s *Server) initHandlers() {
	s.router.Get("/health", s.healthCheckHandler())
	s.router.Post("/profile", s.profileCreateHandler())
	s.router.Post("/profile2", s.profile2CreateHandler())
}

func (s *Server) healthCheckHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log := request.Logger(s.logger, r)
		data, err := s.healthUsecases.Check(r.Context(), log)
		respond.ItemOrFail(w, data, err, log)
	}
}

func (s *Server) profileCreateHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log := request.Logger(s.logger, r)

		dto := profile.CreateDTO{}
		if err := request.DTOFromJSON(&dto, r); err != nil {
			respond.Error(w, err, log)
			return
		}

		data, err := s.profileUsecases.Create(r.Context(), log, dto)

		respond.ItemOrFail(w, data, err, log)
	}
}

func (s *Server) profile2CreateHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		log := request.Logger(s.logger, r)

		dto := profile2.CreateDTO{}
		if err := request.DTOFromJSON(&dto, r); err != nil {
			respond.Error(w, err, log)
			return
		}

		data, err := s.profile2Usecases.Create(r.Context(), log, dto)

		respond.ItemOrFail(w, data, err, log)
	}
}
