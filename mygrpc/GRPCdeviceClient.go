package mygprc

import (
	"os"
	pb_device "github.com/onezerobinary/db-box/proto/device"
	"github.com/goinggo/tracelog"
	"golang.org/x/net/context"
)

func AddDevice (device *pb_device.Device) (response *pb_device.Response) {

	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_device.NewDeviceServiceClient(conn)

	resp, _ := client.AddDevice(context.Background(), device)

	if !resp.Response {
		tracelog.Warning("GRPCdeviceClient", "AddDevice", "Error: Device not added or already present" )
		// Return false as Response
		falseResponse := pb_device.Response{false}
		return &falseResponse
	}

	return resp
}

func GetDeviceByExpoToken (expoPushToken *pb_device.ExpoPushToken) (device *pb_device.Device) {
	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_device.NewDeviceServiceClient(conn)

	device, err := client.GetDeviceByExpoToken(context.Background(), expoPushToken)

	if err != nil {
		tracelog.Errorf(err, "GRPCdeviceClient", "GetDeviceByExpoToken", "Error: Device not retrieved" )
		os.Exit(1)
	}

	return device
}

func UpdateStatus (status *pb_device.Status) (response *pb_device.Response) {
	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)

	client := pb_device.NewDeviceServiceClient(conn)

	resp, err := client.UpdateStatus(context.Background(), status)

	if !resp.Response {
		tracelog.Errorf(err, "GRPCdeviceClient", "UpdateStatus", "Error: Device's status not updated" )
		os.Exit(1)
	}

	return resp
}

func UpdatePosition (position *pb_device.Position) ( response *pb_device.Response) {
	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_device.NewDeviceServiceClient(conn)

	resp, err := client.UpdatePosition(context.Background(), position)

	if !resp.Response {
		tracelog.Errorf(err, "GRPCdeviceClient", "UpdatePosition", "Error: Device's position not updated" )
		os.Exit(1)
	}

	return resp
}

func UpdateMobileNumber (mobileNumber *pb_device.MobileNumber)  (response *pb_device.Response) {
	conn := StartGRPCConnection()
	defer StopGRPCConnection(conn)
	// Search into the DB the user
	client := pb_device.NewDeviceServiceClient(conn)

	resp, err := client.UpdateMobileNumber(context.Background(), mobileNumber)

	if !resp.Response {
		tracelog.Errorf(err, "GRPCdeviceClient", "UpdateMobileNumber", "Error: Device's mobile number not updated" )
		os.Exit(1)
	}

	return resp
}