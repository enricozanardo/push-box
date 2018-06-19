package handler

import (
	"net/http"
	"encoding/json"
	"github.com/onezerobinary/push-box/model"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/onezerobinary/push-box/mygrpc"
	"fmt"
	"github.com/goinggo/tracelog"
	"errors"
)

func TokenHandler(w http.ResponseWriter, req *http.Request) {
	// process form submission
	if req.Method == http.MethodPost {

		decoder := json.NewDecoder(req.Body)

		var data model.MobileUser
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		// Perform all the checks
		token := mygprc.GenerateToken(data.User.Username, data.User.Password)
		accountToken := pb_account.Token{token}

		account := mygprc.GetAccountByToken(&accountToken)

		if account.Username == "" {
			tracelog.Errorf(err, "mobile", "TokenHandler", "Account empty")
			return
		}

		// Add the device to the user if not already present
		expotoken := string(data.Token.Value)

		expoPushTokenDevice := pb_account.ExpoPushToken{expotoken, &accountToken }

		isAdded := mygprc.AddExpoPushToken(&expoPushTokenDevice)

		if !isAdded {
			err := errors.New("Token not added to the account")
			tracelog.Error(err, "mobile", "TokenHandler")
		}

		tracelog.Trace("mobile", "TokenHandler", "Token added to account")
		
		//TODO: Retun "" if no math is found
		//token = ""

		fmt.Fprintf(w,  token)
	}
}
