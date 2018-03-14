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
		User string 		`json:"user"`
	} 						`json:"data"`
}

type MobileUser struct {
	Token struct { Value ExpoPushToken `json:"value"` } `json:"token"`
	User struct{
		Username string `json:"username"`
		Password string `json:"password"`
	} `json:"user"`
}