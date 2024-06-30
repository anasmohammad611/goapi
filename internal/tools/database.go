package tools

import log "github.com/sirupsen/logrus"

type LoginDetails struct {
	AuthToken string
	UserName  string
}

type CoinDetails struct {
	UserName string
	Coins    int64
}

type DatabaseInterface interface {
	GetUserLoginDetails(userName string) (LoginDetails, bool)
	GetUserCoins(userName string) (CoinDetails, bool)
	SetUpDatabase() error
}

func NewDatabase() (DatabaseInterface, error) {
	var database DatabaseInterface = mockDB{}

	err := database.SetUpDatabase()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return database, nil
}
