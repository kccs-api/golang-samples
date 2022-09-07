package structs

// 気象警報データの構造体
type VXWW50ReceivedData struct {
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
	Head                    VXWW50Head     `json:"head"`
	Body                    VXWW50Body     `json:"body"`
	Warning                 VXWW50Warning  `json:"warning"`
}

type VXWW50Head struct {
	TargetDatetime AllowEmptyTime `json:"target_date_time"`
	Serial         string         `json:"serial"`
	Text           string         `json:"text"`
}

type VXWW50Body struct {
	TargetArea VXWW50TargetArea `json:"target_area"`
}

type VXWW50TargetArea struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type VXWW50Warning struct {
	Item []VXWW50Item `json:"item"`
}

type VXWW50Item struct {
	Kind VXWW50Kind `json:"kind"`
	Area VXWW50Area `json:"area"`
}

type VXWW50Kind struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	Status     string `json:"status"`
	StatusCode string `json:"status_code"`
}

type VXWW50Area struct {
	Name string `json:"name"`
	Code string `json:"code"`
}