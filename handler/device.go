package handler

import (
	"net/http"
	"github.com/onezerobinary/push-box/model"
	"encoding/json"
	"fmt"
	pb_device "github.com/onezerobinary/db-box/proto/device"
	"github.com/onezerobinary/push-box/utils"
	"github.com/onezerobinary/push-box/mygrpc"
	"github.com/goinggo/tracelog"
)

func DeviceHandler(w http.ResponseWriter, req *http.Request) {
	// process form submission
	if req.Method == http.MethodPost {

		decoder := json.NewDecoder(req.Body)

		var data model.Device
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		fmt.Println(data.Type)

		// Set the retrieved properties to the device
		device := pb_device.Device{}
		device.Expopushtoken = &pb_device.ExpoPushToken{data.ExpoPushToken}
		device.Type = data.Type
		device.Active = utils.StringToBool(data.Active)
		device.Latitude = utils.StringToFloat32(data.Latitude)
		device.Longitude = utils.StringToFloat32(data.Longitude)
		device.Mobilenumber = data.Mobilenumber

		// Add the device to the BD
		response := mygprc.AddDevice(&device)

		//the callback return a text true/false based on the response
		fmt.Fprintf(w,   utils.BoolToString(response.Response))
	}
}


func PositionHandler(w http.ResponseWriter, req *http.Request) {
	// process position submission
	if req.Method == http.MethodPost {

		decoder := json.NewDecoder(req.Body)

		var data model.Position
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		// Set the retrieved proprieties to the position model
		position := pb_device.Position{}
		position.Expopushtoken = &pb_device.ExpoPushToken{data.ExpoPushToken}
		position.Latitude = utils.StringToFloat32(data.Latitude)
		position.Longitude = utils.StringToFloat32(data.Longitude)

		// Update position
		response := mygprc.UpdatePosition(&position)

		if !response.Response {
			tracelog.Errorf(err, "device", "positionHandler", "It was not possible to update the positon of the device")
			return
		}

		//the callback return a text true/false based on the response
		fmt.Fprintf(w,  utils.BoolToString(response.Response))
	}
}

func StatusHandler(w http.ResponseWriter, req *http.Request){
	// process position submission
	if req.Method == http.MethodPost {

		decoder := json.NewDecoder(req.Body)

		var data model.Status
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		// Set the retrieved proprieties to the position model
		status := pb_device.Status{}
		status.Expopushtoken = &pb_device.ExpoPushToken{data.ExpoPushToken}
		status.Active = utils.StringToBool(data.Active)

		// Update position
		response := mygprc.UpdateStatus(&status)

		if !response.Response {
			tracelog.Errorf(err, "device", "StatusHandler", "It was not possible to update the status of the device")
			return
		}

		//the callback return a text true/false based on the response
		fmt.Fprintf(w,  utils.BoolToString(response.Response))
	}
}


func MobileHandler(w http.ResponseWriter, req *http.Request){
	// process position submission
	if req.Method == http.MethodPost {

		decoder := json.NewDecoder(req.Body)

		var data model.Mobile
		err := decoder.Decode(&data)
		if err != nil {
			panic(err)
		}
		defer req.Body.Close()

		// Set the retrieved proprieties to the position model
		mobile := pb_device.MobileNumber{}
		mobile.Expopushtoken = &pb_device.ExpoPushToken{data.ExpoPushToken}
		mobile.Mobilenumber = data.Mobilenumber

		// Update position
		response := mygprc.UpdateMobileNumber(&mobile)

		if !response.Response {
			tracelog.Errorf(err, "device", "MobileHandler", "It was not possible to update the mobile number of the device")
			return
		}

		//the callback return a text true/false based on the response
		fmt.Fprintf(w,  utils.BoolToString(response.Response))
	}
}

