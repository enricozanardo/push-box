package notification

import (
	"bytes"
	"log"
	"github.com/onezerobinary/push-box/model"
	"encoding/json"
	"net/http"
	"github.com/onezerobinary/push-box/mygrpc"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/goinggo/tracelog"
	"fmt"
)

const (
	APIURL  = "https://exp.host/--/api/v2/push/send"
	NOTIFICATION_LIMIT = 100
)


func SendNotification(tokens []*pb_account.Token){

	notifications := []model.Notification{}

	for _, token := range tokens {

		account, err := mygprc.GetAccountByToken(token)

		if err != nil {
			tracelog.Errorf(err, "expo", "SendNotification", "It was not possible to retrieve the account")
		}

		// Add the device to the user if not already present
		if account.Username == "" {
			tracelog.Errorf(err, "expo", "SendNotification", "Account empty")
			return
		}

		for _, device := range account.Expopushtoken {

			notification := model.Notification{}
			notification.To = model.ExpoPushToken(device)
			notification.Title = "PulseRescue"
			notification.Body = "Emergency"
			notification.Sound = model.SOUND_DEFAULT
			notification.Priority = model.PRIORITY_DEFAULT
			notification.Data.User = account.Username
			notification.Badge = 1

			// Add in the list the notification
			if len(notifications) < NOTIFICATION_LIMIT {
				notifications = append(notifications, notification)
			}
		}
	}

	//Send the messages to all the account devices
	json, err := json.Marshal(notifications)

	client := &http.Client{}
	req, err := http.NewRequest("POST", APIURL, bytes.NewReader(json))

	if err != nil {
		log.Println("error,", err)
	}

	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-encoding", "gzip, deflate")
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	defer resp.Body.Close()

	fmt.Println("Code:", resp.StatusCode)

	if err != nil || resp.StatusCode != 200 {
		tracelog.Error(err, "SendNotification","Error")
	}
}
