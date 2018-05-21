package model

const (
	SOUND_DEFAULT Sound = "default"
	SOUND_NULL    Sound = "null"
	PRIORITY_DEFAULT Priority = "default"
	PRIORITY_NORMAL Priority = "normal"
	PRIORITY_HIGH Priority = "high"
	STATUS_OK PushStatus = "ok"
	STATUS_ERROR PushStatus = "error"
)

type ExpoPushToken string
type Sound string
type Priority string
type PushStatus string

type Notification struct {
	To ExpoPushToken 		`json:"to"`
	Title string 			`json:"title"`
	Body string 			`json:"body"`
	Sound Sound 			`json:"sound"`
	Ttl int32 				`json:"ttl"`
	Expiration int32 		`json:"expiration"`
	Priority Priority 		`json:"priority"`
	Badge int32 			`json:"badge"`
	Data struct{
		Address string 			`json:"address"`
		AddressNumber string 	`json:"addressnumber"`
		PostalCode string 		`json:"postalcode"`
		Place string 			`json:"place"`
		Lat string 				`json:"lat"`
		Lng string 				`json:"lng"`
		Time string 			`json:"time"`
	} 							`json:"data"`
}

type StopNotification struct {
	To ExpoPushToken 		`json:"to"`
	Title string 			`json:"title"`
	Body string 			`json:"body"`
	Sound Sound 			`json:"sound"`
	Ttl int32 				`json:"ttl"`
	Expiration int32 		`json:"expiration"`
	Priority Priority 		`json:"priority"`
	Badge int32 			`json:"badge"`
	Data struct{
		IsActive bool 			`json:"isactive"`
	} 							`json:"data"`
}

type MobileUser struct {
	Token struct { Value ExpoPushToken `json:"value"` } `json:"token"`
	User struct{
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
}

type Device struct {
	ExpoPushToken string 	`json:"expopushtoken"`
	Type string 			`json:"type"`
	Active string 			`json:"active"`
	Latitude string 		`json:"latitude"`
	Longitude string 		`json:"longitude"`
	Mobilenumber string 	`json:"mobilenumber"`
}

type Position struct {
	ExpoPushToken string 	`json:"expopushtoken"`
	Latitude string 		`json:"latitude"`
	Longitude string 		`json:"longitude"`
}

type Status struct {
	ExpoPushToken string 	`json:"expopushtoken"`
	Active string 			`json:"active"`
}

type Mobile struct {
	ExpoPushToken string 	`json:"expopushtoken"`
	Mobilenumber string 	`json:"mobilenumber"`
}