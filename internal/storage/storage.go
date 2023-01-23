package storage

import (
	"burmachine/LinkGenerator/internal/config"
	"burmachine/LinkGenerator/internal/database"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type Storage struct {
	StorageType string
	Postgres    *pgx.Conn
	InMemory    map[string]string
}

func NewStorageStruct(storageType string) *Storage {
	return &Storage{StorageType: storageType}
}

func (s *Storage) StorageInit(conf config.Conf) error {
	if s.StorageType == "postgres" {
		conn, err := database.InitConnection(conf.DbUrl, conf.Addr)
		if err != nil {
			return fmt.Errorf("storage initialization error: %v", err)
		}
		s.Postgres = conn
	} else if s.StorageType == "inmemory" {
		m := make(map[string]string)
		s.InMemory = m
	}
	return nil
}
