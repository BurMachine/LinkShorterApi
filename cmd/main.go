package main

import (
	"flag"
	"fmt"
	"log"
	"sync"

	"burmachine/LinkGenerator/internal/config"
	grpcHandlers2 "burmachine/LinkGenerator/internal/handlers/grpc"
	httpHandlers "burmachine/LinkGenerator/internal/handlers/http"
	server2 "burmachine/LinkGenerator/internal/server"
	"burmachine/LinkGenerator/internal/storage"
	"burmachine/LinkGenerator/pkg/flagsHandling"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	storageTypeFlag := flag.String("storage", "postgres", "Storage selection (have postgres or inmemory values")
	cfgPath := flag.String("config", "./config.yaml", "Path to yaml configuration file")
	flag.Parse()

	// Инициализация конфигурации
	conf := config.NewConfigStruct()
	err := conf.LoadConfig(*cfgPath)
	if err != nil {
		log.Fatalln(err)
	}

	// Инициализация хранилища
	storageType, err := flagsHandling.CheckFlagString(*storageTypeFlag)
	if err != nil {
		log.Fatalln(err)
	}
	var storageS storage.ServiceStorage
	if storageType == "inmemory" {
		storageS = storage.NewInMemoryStorageInit()
	} else if storageType == "postgres" {
		storageS, err = storage.NewStorageInit(*conf)
		if err != nil {
			log.Fatalln(err)
		}
	}
	grpcHandlers := &grpcHandlers2.GrpcHandlers{Storage: &storageS}
	grpcServ := grpc.NewServer()
	server := server2.NewServerWithConfiguration(*conf)
	server.GrpcHadles = grpcHandlers
	server.GrpcServ = grpcServ
	mux := runtime.NewServeMux()
	server.Mux = mux

	var handlersHttp httpHandlers.HttpHandlers
	handlersHttp.Storage = &storageS

	err = mux.HandlePath("POST", "/generate", handlersHttp.GenerateShortLink)
	if err != nil {
		err = fmt.Errorf("handler registration error: %v", err)
	}
	err = mux.HandlePath("GET", "/", handlersHttp.GetOriginalUrl)
	if err != nil {
		err = fmt.Errorf("handler registration error: %v", err)
	}

	wg := &sync.WaitGroup{}
	wg.Add(2)
	errChan := make(chan error)
	server.ErrorChan = errChan

	go func(chan error) {
		select {
		case <-errChan:
			log.Println("[SERVER] - running error")
		}
	}(errChan)
	ctx := context.Background()
	err = server.Run(ctx, wg)
	if err != nil {
		log.Fatalln(err)
	}
	wg.Wait()
}
