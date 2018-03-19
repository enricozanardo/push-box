package notification

import (
	"bytes"
	"log"
	"github.com/onezerobinary/push-box/model"
	"net/http"
	"github.com/onezerobinary/push-box/mygrpc"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/goinggo/tracelog"
	"errors"
	"fmt"
	"encoding/json"
	"io/ioutil"
)

const (
	APIURL  = "https://exp.host/--/api/v2/push/send"
	NOTIFICATION_LIMIT = 100
)


func SendNotification(tokens []*pb_account.Token) (statusCode *int, err error){

	notifications := []model.Notification{}

	for _, token := range tokens {

		account, err := mygprc.GetAccountByToken(token)

		var statusCode int

		if err != nil {
			tracelog.Errorf(err, "expo", "SendNotification", "It was not possible to retrieve the account")
			statusCode = 400
			return &statusCode, err
		}

		// Add the device to the user if not already present
		if account.Username == "" {
			err = errors.New("Error: Account empty")
			tracelog.Errorf(err, "expo", "SendNotification", "Error: Account empty")
			statusCode = 400
			return &statusCode, err
		}

		for _, device := range account.Expopushtoken {

			notification := model.Notification{}
			notification.To = model.ExpoPushToken(device)
			notification.Title = "Io so dea Lazio"
			notification.Body = "Mica do Frosinone"
			notification.Sound = model.SOUND_DEFAULT
			notification.Priority = model.PRIORITY_DEFAULT
			//TODO: add data of the emergency
			notification.Data.User = account.Username
			notification.Badge = 0

			// Add in the list the notification
			if len(notifications) < NOTIFICATION_LIMIT {
				notifications = append(notifications, notification)
			}
		}
	}

	//Send the messages to all the account devices
	jsonData, err := json.Marshal(notifications)

	client := &http.Client{}
	req, err := http.NewRequest("POST", APIURL, bytes.NewReader(jsonData))

	if err != nil {
		log.Println("error,", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()


	//type Callback struct {
	//		Status string 		`json:"status"`
	//		Message string		`json:"message"`
	//		Details struct{
	//			Error string 	`json:"error"`
	//			Sns struct{
	//				StatusCode string 	`json:"statusCode"`
	//				Reason string 	`json:"reason"`
	//				Message string 	`json:"__message"`
	//			} 					`json:"sns"`
	//		} 					`json:"details"`
	//}

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))

	return &resp.StatusCode, nil
}
