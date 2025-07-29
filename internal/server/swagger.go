package server

import httpSwagger "github.com/swaggo/http-swagger/v2"

func (s *Server) initSwagger() {
	s.logger.Info("Init swagger...")

	s.router.Mount("/swagger", httpSwagger.WrapHandler)

	// s.router.Get("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL(s.config.Rest.Listen+"/swagger/doc.json"),
	// ))
}
