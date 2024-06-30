package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/anasmohammad611/goapi/api"
	"github.com/anasmohammad611/goapi/internal/tools"
	"github.com/gorilla/schema"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func GetCoinBalance(w http.ResponseWriter, r *http.Request) {
	var params api.CoinBalanceParams

	decoder := schema.NewDecoder()
	err := decoder.Decode(&params, r.URL.Query())

	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w)
		return
	}

	database, err := tools.NewDatabase()
	if err != nil {
		api.InternalServerErrorHandler(w)
		return
	}

	fmt.Println(`Fetching coins`)
	tokenDetails, flag := database.GetUserCoins(params.Username)
	if flag == false {
		log.Error(err)
		api.InternalServerErrorHandler(w)
		return
	}

	response := api.CoinBalanceRes{
		Balance: tokenDetails.Coins,
		Code:    http.StatusOK,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)

	if err != nil {
		log.Error(err)
		api.InternalServerErrorHandler(w)
		return
	}
}
