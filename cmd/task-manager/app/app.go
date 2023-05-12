package app

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net"
	v1 "rohan.com/task-manager/apis/v1/generated"
	config "rohan.com/task-manager/internal/config"
	v1user "rohan.com/task-manager/pkg/task-manager/handlers/v1/user"
	"rohan.com/task-manager/pkg/task-manager/repo/postgres"
	"rohan.com/task-manager/pkg/task-manager/repo/postgres/bun"
)

type app struct {
	dbService  postgres.DBService
	grpcServer *grpc.Server
	ctx        context.Context
}

func New() *app {
	ctx := context.Background()
	return &app{
		ctx: ctx,
	}
}

func (a *app) Start() error {
	fmt.Println("connecting to db ....")
	a.dbService = a.createDBConnection()
	fmt.Println("connected to db ....")
	fmt.Println("creating tables ....")
	err := a.createDBtables()
	fmt.Println("created tables ....")
	if err != nil {
		fmt.Println("error creating db tables")
		return err
	}
	flag.Parse()
	listener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 9091))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	fmt.Println("registering services...")
	//register user service
	userService := v1user.NewService(a.dbService, a.ctx)
	v1.RegisterUserServiceServer(grpcServer, userService)
	fmt.Println("registered services...")
	// register task service
	//v1.RegisterTaskServiceServer(grpcServer)
	reflection.Register(grpcServer)
	a.grpcServer = grpcServer
	fmt.Println("grpc listening on 9091...")
	err = grpcServer.Serve(listener)
	if err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return err
	}
	return nil
}

func (a *app) createDBtables() error {
	err := a.dbService.CreateUserTable(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (a *app) createDBConnection() *bun.DBClient {
	// read config
	data, err := ioutil.ReadFile("../../internal/config/config.yaml")
	if err != nil {
		log.Fatal(err)
	}

	// unmarshal the YAML data into a Config struct
	var dbConfig config.Config
	err = yaml.Unmarshal(data, &dbConfig)
	if err != nil {
		log.Fatal(err)
	}
	pgHostname := dbConfig.Host
	pgUsername := dbConfig.User
	pgPwd := dbConfig.Password
	pgPort := dbConfig.Port
	pgDB := dbConfig.DBName

	serverString := fmt.Sprintf("postgres://%s:%s@%s:%d/", pgUsername, pgPwd, pgHostname, pgPort)
	opts := "?sslmode=disable"

	var db *bun.DBClient
	var i int
	for i = 0; i < 3; i++ {
		db, err = bun.InitDatabase(serverString, pgDB, opts)
		if db == nil || err != nil {
			fmt.Println("retrying db connection....")
		} else {
			fmt.Println("connected to db ....")
			break
		}
	}

	return db

}

func (a *app) Shutdown() error {
	a.grpcServer.GracefulStop()
	a.grpcServer.Stop()
	err := a.dbService.Close()
	if err != nil {
		return err
	}
	return nil
}
