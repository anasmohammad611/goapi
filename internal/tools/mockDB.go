package tools

import "time"

type mockDB struct {
}

var mockLoginDetails = map[string]LoginDetails{
	"alex": {
		AuthToken: "123ABC",
		UserName:  "alex",
	},
	"john": {
		AuthToken: "123ABC",
		UserName:  "john",
	},
	"kale": {
		AuthToken: "123ABC",
		UserName:  "kale",
	},
}

var mockCoinDetails = map[string]CoinDetails{
	"alex": {
		Coins:    300,
		UserName: "alex",
	},
	"joh": {
		Coins:    1200,
		UserName: "john",
	},
	"kale": {
		Coins:    10,
		UserName: "kale",
	},
}

func (d mockDB) GetUserLoginDetails(username string) (LoginDetails, bool) {
	time.Sleep(1 * time.Second)

	var clientData = LoginDetails{}
	clientData, ok := mockLoginDetails[username]
	if !ok {
		return LoginDetails{}, false
	}

	return clientData, true
}

func (d mockDB) GetUserCoins(username string) (CoinDetails, bool) {
	time.Sleep(1 * time.Second)

	var clientData = CoinDetails{}
	clientData, ok := mockCoinDetails[username]
	if !ok {
		return CoinDetails{}, false
	}

	return clientData, true
}

func (d mockDB) SetUpDatabase() error {
	return nil
}
