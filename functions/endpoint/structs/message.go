package structs

import "google.golang.org/api/pubsub/v1"

// Pubsubから受信するデータの構造体
type PushRequest struct {
	Message      pubsub.PubsubMessage `json:"message"`
	Subscription string               `json:"subscription"`
}

// 災害データの構造体
type ReceivedData struct {
	ID                      string         `json:"id"`
	DataTypeCode            string         `json:"data_type_code"`
	ReportControlStatus     string         `json:"report_control_status"`
	ReportControlStatusCode string         `json:"report_control_status_code"`
	PublishingOffice        string         `json:"publishing_office"`
	EventID                 string         `json:"event_id"`
	Region                  string         `json:"region"`
	ReportDatetime          AllowEmptyTime `json:"report_datetime"`
	ReceivedDatetime        AllowEmptyTime `json:"received_datetime"`
	RegisterDatetime        AllowEmptyTime `json:"register_datetime"`
	InfoType                string         `json:"info_datetime"`
	Head                    Head           `json:"head"`
	Body                    Body           `json:"body"`
	Warning                 Warning        `json:"warning"`
}

type Head struct {
	TargetDatetime AllowEmptyTime `json:"target_date_time"`
	Serial         string         `json:"serial"`
	Text           string         `json:"text"`
}

type Body struct {
	TargetArea TargetArea `json:"target_area"`
}

type TargetArea struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type Warning struct {
	Item []Item `json:"item"`
}

type Item struct {
	Kind Kind `json:"kind"`
	Area Area `json:"area"`
}

type Kind struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}

type Area struct {
	Name string `json:"name"`
	Code string `json:"code"`
}