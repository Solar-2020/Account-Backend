package main

import (
	"database/sql"
	"github.com/Solar-2020/Account-Backend/cmd/handlers"
	accountHandler "github.com/Solar-2020/Account-Backend/cmd/handlers/account"
	"github.com/Solar-2020/Account-Backend/internal/clients/auth"
	"github.com/Solar-2020/Account-Backend/internal/services/account"
	"github.com/Solar-2020/Account-Backend/internal/storages/accountStorage"
	authapi "github.com/Solar-2020/Authorization-Backend/pkg/api"
	"github.com/Solar-2020/GoUtils/common"
	"github.com/Solar-2020/GoUtils/context/session"
	"github.com/Solar-2020/GoUtils/http/errorWorker"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"os"
	"os/signal"
	"syscall"
)

type config struct {
	common.SharedConfig
	AccountDataBaseConnectionString string `envconfig:"ACCOUNT_DB_CONNECTION_STRING" default:"-"`
	ServerSecret                    string `envconfig:"SERVER_SECRET" default:"Basic secret"`
}

func main() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	authorizationDB, err := sql.Open("postgres", cfg.AccountDataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	authorizationDB.SetMaxIdleConns(5)
	authorizationDB.SetMaxOpenConns(10)

	errorWorker := errorWorker.NewErrorWorker()

	accountStorage := accountStorage.NewStorage(authorizationDB)
	accountService := account.NewService(accountStorage)
	accountTransport := account.NewTransport()

	accountHandler := accountHandler.NewHandler(accountService, accountTransport, errorWorker)
	authService := authapi.AuthClient{
		Addr: cfg.AuthServiceAddress,
	}
	session.RegisterAuthService(&authService)

	authClient := auth.NewClient(cfg.AuthServiceAddress, cfg.ServerSecret)
	middlewares := handlers.NewMiddleware(&log, authClient)

	server := fasthttp.Server{
		Handler: handlers.NewFastHttpRouter(accountHandler, middlewares).Handler,
	}

	go func() {
		log.Info().Str("msg", "start server").Str("port", cfg.Port).Send()
		if err := server.ListenAndServe(":" + cfg.Port); err != nil {
			log.Error().Str("msg", "server run failure").Err(err).Send()
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	defer func(sig os.Signal) {

		log.Info().Str("msg", "received signal, exiting").Str("signal", sig.String()).Send()

		if err := server.Shutdown(); err != nil {
			log.Error().Str("msg", "server shutdown failure").Err(err).Send()
		}

		//dbConnection.Shutdown()
		log.Info().Str("msg", "goodbye").Send()
	}(<-c)
}

//func grpcListener(accountService account.Service) {
//	listener, err := net.Listen("tcp", ":5300")
//
//	if err != nil {
//		grpclog.Fatalf("failed to listen: %v", err)
//	}
//
//	opts := []grpc.ServerOption{}
//	grpcServer := grpc.NewServer(opts...)
//
//	accountpb.RegisterAccountServer(grpcServer, accountService)
//	grpcServer.Serve(listener)
//}
