package tools

import (
	log "github.com/sirupsen/logrus"
)

type LoginDetails struct {
	AuthToken string
	Username  string
}

type PointDetails struct {
	Username string
	Balance  int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(username string) *LoginDetails
	GetUserPointDetails(username string) *PointDetails
	SetupDatabase() error
}

func NewDatabase() (*DatabaseInterface, error) {
	var database DatabaseInterface = &mockDatabase{}

	var err error = database.SetupDatabase()

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &database, nil
}
