package notification

import (
	"testing"
	pb_account "github.com/onezerobinary/db-box/proto/account"
	"github.com/goinggo/tracelog"
)

func TestSendNotification(t *testing.T) {

	tracelog.Start(tracelog.LevelTrace)
	defer tracelog.Stop()

	tokens := []*pb_account.Token{}

	token := pb_account.Token{"2284fe70432bbef5a5354653c88d8e5cda2880dd"}

	tokens = append(tokens, &token)

	SendNotification(tokens)
}