package main

import (
	"burmachine/LinkGenerator/internal/config"
	httpHandlers "burmachine/LinkGenerator/internal/handlers/http"
	server2 "burmachine/LinkGenerator/internal/server"
	"burmachine/LinkGenerator/internal/storage"
	"burmachine/LinkGenerator/pkg/flagsHandling"
	"flag"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
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

	server := server2.NewServerWithConfiguration(*conf)
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

	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
