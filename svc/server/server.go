package server

import (
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"test-fizz-buzz/svc/configs"
	"test-fizz-buzz/svc/rest"
)

type Server struct {
	logger     logrus.FieldLogger
	restServer *rest.Server
}

func NewServer(c *configs.Config) (*Server, error) {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)

	restServer, err := rest.NewServer(logger, c)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create a new rest server")
	}

	return &Server{
		logger:     logger,
		restServer: restServer,
	}, nil
}

func (s *Server) Run() error {

	if s.restServer != nil {
		err := s.restServer.Run()
		if err != nil {
			return err
		}
	} else {
		s.logger.Infoln("no rest server started")
	}

	return nil
}
