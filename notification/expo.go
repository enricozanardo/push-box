package notification

import (
	"bytes"
	"log"
	"github.com/onezerobinary/push-box/model"
	"net/http"
	"fmt"
	"encoding/json"
	"io/ioutil"
	pb_push "github.com/onezerobinary/push-box/proto"
)

const (
	APIURL  = "https://exp.host/--/api/v2/push/send"
	NOTIFICATION_LIMIT = 100
)

func StopNotifications (stop *pb_push.Stop) (stopResponse *pb_push.StopResponse, err error) {

	notifications := []model.StopNotification{}

	response := pb_push.StopResponse{}

	for _, device := range stop.DeviceTokens {

		notification := model.StopNotification{}
		notification.To = model.ExpoPushToken(device)
		notification.Title = "PulseRescue"
		notification.Body = "Emergency Closed"
		notification.Sound = model.SOUND_DEFAULT
		notification.Priority = model.PRIORITY_DEFAULT
		//add data of the notification
		notification.Data.IsActive = stop.IsActive
		notification.Badge = 0

		// Add in the list the notification
		if len(notifications) < NOTIFICATION_LIMIT {
			notifications = append(notifications, notification)
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

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))


	response.IsClosed = true
	// Send back the response
	stopResponse = &response

	return stopResponse, nil
}


func SendNotification(info *pb_push.Info) (statusCode *int, err error){

	notifications := []model.Notification{}

	//for _, token := range info.DeviceTokens {
	//
	//	accountToken := pb_account.Token{token}
	//	account, err := mygprc.GetAccountByToken(&accountToken)
	//
	//	var statusCode int
	//
	//	if err != nil {
	//		tracelog.Errorf(err, "expo", "SendNotification", "It was not possible to retrieve the account")
	//		statusCode = 400
	//		return &statusCode, err
	//	}
	//
	//	// Add the device to the user if not already present
	//	if account.Username == "" {
	//		err = errors.New("Error: Account empty")
	//		tracelog.Errorf(err, "expo", "SendNotification", "Error: Account empty")
	//		statusCode = 400
	//		return &statusCode, err
	//	}

	// Android -> ExponentPushToken[VqalPOCUT5DVmVUpf6Qq3B]
	// iphone -> ExponentPushToken[APcW1mOPoMiiB_apLgu5PS]


	for _, device := range info.DeviceTokens {

		notification := model.Notification{}
		notification.To = model.ExpoPushToken(device)
		notification.Title = "Emergency"
		notification.Body = info.Emergency.Address + " " + info.Emergency.AddressNumber
		notification.Sound = model.SOUND_DEFAULT
		notification.Priority = model.PRIORITY_DEFAULT
		//add data of the emergency
		notification.Data.Address = info.Emergency.Address
		notification.Data.AddressNumber = info.Emergency.AddressNumber
		notification.Data.PostalCode = info.Emergency.PostalCode
		notification.Data.Place = info.Emergency.Place
		notification.Data.Lat = info.Emergency.Lat
		notification.Data.Lng = info.Emergency.Lng
		notification.Data.Time = info.Emergency.Time
		notification.Badge = 0

		// Add in the list the notification
		if len(notifications) < NOTIFICATION_LIMIT {
			notifications = append(notifications, notification)
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

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(string(body))

	return &resp.StatusCode, nil
}
