package postgres

import (
	"burmachine/LinkGenerator/internal/config"
	"burmachine/LinkGenerator/internal/database"
	"fmt"
	"github.com/jackc/pgx/v5"
)

type dbStorage struct {
	Conn *pgx.Conn
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

func (s *dbStorage) GenerateShortLink() {
	//TODO implement me
	panic("implement me")
}

func (s *dbStorage) GetFullLink() {
	//TODO implement me
	panic("implement me")
}
