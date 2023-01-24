package storage

import (
	"burmachine/LinkGenerator/internal/config"
	"burmachine/LinkGenerator/internal/database"
	"burmachine/LinkGenerator/internal/serviceErrors"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type dbStorage struct {
	Conn *pgx.Conn
}

func (s *dbStorage) AddShortLink(link, shortLink string) error {
	err := s.CheckLinksDb(link, shortLink, s.Conn)
	if err != nil {
		return err
	}
	err = database.Insert(link, shortLink, s.Conn)
	if err != nil {
		return err
	}
	return nil
}

func (s *dbStorage) GetFullLink(shortLink string) (string, error) {
	original := database.SelectWhereShortLinkIs(shortLink, s.Conn)
	if original == "" {
		return "", serviceErrors.ErrorLinkNotExist
	}
	return original, nil
}

func NewStorageInit(conf config.Conf) (*dbStorage, error) {
	s := dbStorage{}
	conn, err := database.InitConnection(conf.DbUrl)
	if err != nil {
		return &s, fmt.Errorf("storage initialization error: %v", err)
	}
	s.Conn = conn
	return &s, nil
}

func (s *dbStorage) CheckLinksDb(link string, short string, conn *pgx.Conn) error {
	original, short := database.SelectWhereLinkIs(link, conn)
	if original != "" {
		return serviceErrors.ErrorLinkExist
	}
	return nil
}
