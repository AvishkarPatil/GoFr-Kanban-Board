package services

import (
	"database/sql"
	"fmt"
)

type DatabaseService struct {
	db *sql.DB
}

func NewDatabaseService(db *sql.DB) *DatabaseService {
	return &DatabaseService{
		db: db,
	}
}

func (ds *DatabaseService) TestConnection() error {
	_, err := ds.db.Query("SELECT 1")
	if err != nil {
		return fmt.Errorf("database connection failed: %w", err)
	}
	return nil
}

func (ds *DatabaseService) GetDB() *sql.DB {
	return ds.db
}