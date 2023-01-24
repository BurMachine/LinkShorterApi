package main

import (
	"burmachine/LinkGenerator/internal/config"
	"burmachine/LinkGenerator/internal/interfaces"
	server2 "burmachine/LinkGenerator/internal/server"
	"burmachine/LinkGenerator/internal/storage/memory"
	"burmachine/LinkGenerator/internal/storage/postgres"
	"burmachine/LinkGenerator/pkg/flagsHandling"
	"flag"
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
	var storage interfaces.Storage
	if storageType == "inmemory" {
		storage = memory.NewStorageInit()
	} else if storageType == "postgres" {
		storage, err = postgres.NewStorageInit(*conf)
		if err != nil {
			log.Fatalln(err)
		}
	}
	println(&storage)

	server := server2.NewServerWithConfiguration(*conf)
	err = server.Run()
	if err != nil {
		log.Fatalln(err)
	}
}
