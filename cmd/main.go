package main

import (
	"database/sql"
	"github.com/Solar-2020/Account-Backend/cmd/handlers"
	accountHandler "github.com/Solar-2020/Account-Backend/cmd/handlers/account"
	"github.com/Solar-2020/Account-Backend/internal/errorWorker"
	accountpb "github.com/Solar-2020/Account-Backend/internal/proto"
	"github.com/Solar-2020/Account-Backend/internal/services/account"
	"github.com/Solar-2020/Account-Backend/internal/storages/accountStorage"
	"github.com/kelseyhightower/envconfig"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/valyala/fasthttp"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type config struct {
	Port                          string `envconfig:"PORT" default:"8099"`
	AuthorizationDataBaseConnectionString string `envconfig:"AUTHORIZATION_DB_CONNECTION_STRING" default:"-"`
}

func main() {
	log := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})

	var cfg config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}


	authorizationDB, err := sql.Open("postgres", cfg.AuthorizationDataBaseConnectionString)
	if err != nil {
		log.Fatal().Msg(err.Error())
		return
	}

	authorizationDB.SetMaxIdleConns(5)
	authorizationDB.SetMaxOpenConns(10)

	errorWorker := errorWorker.NewErrorWorker()

	authorizationStorage := accountStorage.NewStorage(authorizationDB)
	authorizationService := account.NewService(authorizationStorage)
	authorizationTransport := account.NewTransport()

	authorizationHandler := accountHandler.NewHandler(authorizationService, authorizationTransport, errorWorker)

	middlewares := handlers.NewMiddleware()

	server := fasthttp.Server{
		Handler: handlers.NewFastHttpRouter(authorizationHandler, middlewares).Handler,
	}



	listener, err := net.Listen("tcp", ":5300")

	if err != nil {
		grpclog.Fatalf("failed to listen: %v", err)
	}

	opts := []grpc.ServerOption{}
	grpcServer := grpc.NewServer(opts...)

	accountpb.RegisterAccountServer(grpcServer, authorizationService)
	grpcServer.Serve(listener)





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
