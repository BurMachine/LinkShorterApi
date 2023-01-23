package main

import (
	"burmachine/LinkGenerator/internal/config"
	myStorage "burmachine/LinkGenerator/internal/storage"
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
	storage := myStorage.NewStorageStruct(storageType)
	err = storage.StorageInit(*conf)
	if err != nil {
		log.Fatalln(err)
	}

}
