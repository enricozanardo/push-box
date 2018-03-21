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
	fakeEmegency := pb_push.Emergency{}
	fakeEmegency.Address = "Via Roma"
	fakeEmegency.AddressNumber = "42"
	fakeEmegency.PostalCode = "39100"
	fakeEmegency.Place = "Bolzano"
	fakeEmegency.Lat = "46.4894107"
	fakeEmegency.Lng = "11.3208888"
	fakeEmegency.Time = "2018-03-21T09:47:42.140Z"

	info := pb_push.Info{}
	info.Emergency = &fakeEmegency

	token := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}

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