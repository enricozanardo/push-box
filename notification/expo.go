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
)

const APIURL  = "https://exp.host/--/api/v2/push/send"


func SendNotification(token *pb_account.Token){

	account, err := mygprc.GetAccountByToken(token)

	if err != nil {
		tracelog.Errorf(err, "expo", "SendNotification", "It was not possible to retrieve the account")
	}

	// Add the device to the user if not already present
	if account.Username == "" {
		tracelog.Errorf(err, "expo", "SendNotification", "Account empty")
		return
	}

	notifications := []model.Notification{}

	//Send the messages to all the account devices
	for _, device := range account.Expopushtoken {

		notification := model.Notification{}
		notification.To = model.ExpoPushToken(device)
		notification.Title = "PulseRescue"
		notification.Body = "Emergency"
		notification.Sound = model.SOUND_DEFAULT
		notification.Priority = model.PRIORITY_DEFAULT
		notification.Data.User = account.Username
		notification.Badge = 0

		// Add in the list the notification
		notifications = append(notifications, notification)
	}

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

	if resp.StatusCode != 200 {
		//Error
	}

	if err != nil {
		log.Println("error 2", err)
	}

	defer resp.Body.Close()
}
