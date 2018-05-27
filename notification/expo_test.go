package notification

import (
	"testing"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/goinggo/tracelog"
	"fmt"
	pb_push "github.com/onezerobinary/push-box/proto"
	"github.com/onezerobinary/push-box/mygrpc"
)

func TestSendNotification(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//Fake emergency
	fakeEmergency := pb_push.Emergency{}
	fakeEmergency.Address = "Triq il San Pawl"
	fakeEmergency.AddressNumber = "396"
	fakeEmergency.PostalCode = "39100"
	fakeEmergency.Place = "San Paul il-Bahar"
	fakeEmergency.Lat = "35.948621"
	fakeEmergency.Lng = "14.399897"
	fakeEmergency.Time = "2018-03-21T09:47:42.140Z"
	fakeEmergency.IsActive = true

	info := pb_push.Info{}
	info.Emergency = &fakeEmergency

	//TODO: Search and add to the info only the devices that are near to the emergency
	//TODO: Store them into an array and when the emergency is closed send a notification to close the emergency
	
	//token := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}
	token := pb_account.Token{"d0a1a743194ff28f049f47b9b69c51563c2cfadf"} // local
	//token := pb_account.Token{"46a249c795cda18c1d8143a781871e1e95d2e011"} //remote

	fakeAccount, err := mygprc.GetAccountByToken(&token)

	if err != nil {
		tracelog.Error(err, "expo_test", "TestSendNotification")
	}

	for _, device := range fakeAccount.Expopushtoken {
		info.DeviceTokens = append(info.DeviceTokens, device)
	}

	status, err := SendNotification(&info)

	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("code: ", &status)
}

func TestSendStopNotification(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	//Fake emergency
	fakeEmergency := pb_push.Emergency{}
	fakeEmergency.Address = "Triq il San Pawl"
	fakeEmergency.AddressNumber = "396"
	fakeEmergency.PostalCode = "39100"
	fakeEmergency.Place = "San Paul il-Bahar"
	fakeEmergency.Lat = "35.948621"
	fakeEmergency.Lng = "14.399897"
	fakeEmergency.Time = "2018-03-21T09:47:42.140Z"
	fakeEmergency.IsActive = false

	info := pb_push.Info{}
	info.Emergency = &fakeEmergency

	//TODO: Search and add to the info only the devices that are near to the emergency
	//TODO: Store them into an array and when the emergency is closed send a notification to close the emergency

	//TODO: use the geo-box


	//token := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}
	token := pb_account.Token{"d0a1a743194ff28f049f47b9b69c51563c2cfadf"} // local
	//token := pb_account.Token{"46a249c795cda18c1d8143a781871e1e95d2e011"} //remote

	fakeAccount, err := mygprc.GetAccountByToken(&token)

	if err != nil {
		tracelog.Error(err, "expo_test", "TestSendNotification")
	}

	for _, device := range fakeAccount.Expopushtoken {
		info.DeviceTokens = append(info.DeviceTokens, device)
	}

	status, err := SendNotification(&info)

	if err != nil {
		fmt.Println("err: ", err)
	}

	fmt.Println("code: ", &status)
}
