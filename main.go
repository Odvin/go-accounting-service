package main

import (
	"context"
	"log"
	"net"

	"github.com/Odvin/go-accounting-service/api"
	db "github.com/Odvin/go-accounting-service/db/sqlc"
	"github.com/Odvin/go-accounting-service/gapi"
	"github.com/Odvin/go-accounting-service/pb"
	"github.com/Odvin/go-accounting-service/util"
	"github.com/jackc/pgx/v5/pgxpool"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)

	// go runGinServer(config, store)
	runGrpcServer(config, store)
}

func runGrpcServer(config util.Config, store db.Store) {

	server, err := gapi.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create gRPC server:", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAccountingServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GrpcServerAddress)
	if err != nil {
		log.Fatal("cannot create listener", err)
	}

	log.Printf("start gRPC server at %s", listener.Addr().String())

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("cannot start gRPC server:", err)
	}
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal("cannot create HTTP server:", err)
	}

	err = server.Start(config.HttpServerAddress)

	if err != nil {
		log.Fatal("cannot start HTTP server:", err)
	}
}
