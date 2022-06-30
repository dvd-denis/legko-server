package apiserver

import (
	"net/http"

	"github.com/dvd-denis/legko-server/internal/app/handler"
	"github.com/dvd-denis/legko-server/internal/app/store"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// APIServer ...
type APIServer struct {
	config  *Config
	logger  *logrus.Logger
	handler *handler.Handler
	store   *store.Store
	router  *gin.Engine
}

// New ...
func New(config *Config) *APIServer {
	return &APIServer{
		config: config,
		logger: logrus.New(),
	}
}

// Start ...
func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	if err := s.ConfigureStore(); err != nil {
		return err
	}

	s.ConfigureRouter()

	s.logger.Info("starting api server")

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

// Configure logger level
func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) ConfigureStore() error {
	st := store.New(s.config.Store)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIServer) ConfigureRouter() {
	s.handler = handler.New(s.store)
	s.router = s.handler.InitRouter()
}
