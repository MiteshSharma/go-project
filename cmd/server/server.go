package server

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/MiteshSharma/project/bi"
	"github.com/MiteshSharma/project/eventdispatcher"

	"github.com/urfave/negroni"

	"github.com/MiteshSharma/project/api"
	"github.com/MiteshSharma/project/app"
	"github.com/MiteshSharma/project/logger"
	"github.com/MiteshSharma/project/metrics"
	"github.com/MiteshSharma/project/middleware"
	"github.com/MiteshSharma/project/model"
	"github.com/MiteshSharma/project/repository"
	"github.com/MiteshSharma/project/setting"
	"github.com/gorilla/mux"
)

type Server struct {
	API        *api.API
	App        *app.App
	Repository repository.Repository
	Router     *mux.Router
	Config     *model.Config
	Metrics    metrics.Metrics
	Log        logger.Logger
	httpServer *http.Server
}

func NewServer(settingData *setting.Setting) *Server {
	config := setting.GetConfig()
	logger := logger.NewLogger(config)
	metrics := metrics.NewMetrics()
	router := mux.NewRouter()
	repository := repository.NewPersistentCacheRepository(logger, config, metrics)
	eventDispatcher := eventdispatcher.NewEventDispatcher(logger, 10, 2)
	biEventHandler := bi.NewBiEventHandler(eventDispatcher)

	appOption := &app.AppOption{
		Config:         config,
		Setting:        settingData,
		Log:            logger,
		Metrics:        metrics,
		Repository:     repository,
		BiEventHandler: biEventHandler,
	}

	api := api.NewAPI(router, appOption, config, metrics, logger)
	server := &Server{
		API:        api,
		Log:        logger,
		Metrics:    metrics,
		Config:     config,
		Router:     router,
		Repository: repository,
	}

	return server
}

func (s *Server) StartServer() {
	n := negroni.New()
	n.UseFunc(middleware.NewLoggerMiddleware(s.Log).GetMiddlewareHandler())
	if s.Config.ZipkinConfig.IsEnable {
		n.UseFunc(middleware.NewZipkinMiddleware(s.Log, "project", s.Config.ZipkinConfig).GetMiddlewareHandler())
	}

	n.UseHandler(s.Router)

	listenAddr := (":" + s.Config.ServerConfig.Port)
	s.Log.Debug("Staring server", logger.String("address", listenAddr))
	s.httpServer = &http.Server{
		Handler:      n,
		Addr:         listenAddr,
		ReadTimeout:  s.Config.ServerConfig.ReadTimeout * time.Second,
		WriteTimeout: s.Config.ServerConfig.WriteTimeout * time.Second,
	}

	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			s.Log.Error("Error starting server ", logger.Error(err))
			return
		}
	}()
}

func (s *Server) StopServer() {
	s.Repository.Close()
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	s.httpServer.Shutdown(ctx)

	os.Exit(0)
}
