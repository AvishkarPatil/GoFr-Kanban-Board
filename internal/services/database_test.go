package services

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestDatabaseService_TestConnection_Success(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT 1").WillReturnRows(sqlmock.NewRows([]string{"1"}).AddRow(1))
	
	service := NewDatabaseService(db)
	err = service.TestConnection()
	
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDatabaseService_TestConnection_Error(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectQuery("SELECT 1").WillReturnError(assert.AnError)
	
	service := NewDatabaseService(db)
	err = service.TestConnection()
	
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "database connection failed")
	assert.NoError(t, mock.ExpectationsWereMet())
}