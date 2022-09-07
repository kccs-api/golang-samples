package structs

type VXSE53ReceivedData struct {
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
	Headline                VXSE53Headline `json:"headline"`
	Body                    VXSE53Body     `json:"body"`
}

type VXSE53Headline struct {
	Text        string              `json:"text"`
	Information []VXSE53Information `json:"information"`
}

type VXSE53Information struct {
	Item []VXSE53Item `json:"item"`
}

type VXSE53Item struct {
	Kind  VXSE53Kind  `json:"kind"`
	Areas VXSE53Areas `json:"areas"`
}

type VXSE53Kind struct {
	Name string `json:"name"`
}

type VXSE53Areas struct {
	Area []VXSE53Area `json:"area"`
}

type VXSE53Area struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

type VXSE53Body struct {
	EarthQuake VXSE53EarthQuake `json:"earthquake"`
	Intensity  VXSE53Intensity  `json:"intensity"`
	Text       string           `json:"text"`
	Comments   VXSE53Comments   `json:"comments"`
}

type VXSE53EarthQuake struct {
	OriginTime           AllowEmptyTime   `json:"origin_time"`
	ArrivalTime          AllowEmptyTime   `json:"arrival_time"`
	Hypocenter           VXSE53Hypocenter `json:"hypocenter"`
	Magnitude            string           `json:"magnitude"`
	MagnitudeDescription string           `json:"magnitude_description"`
}

type VXSE53Hypocenter struct {
	Area   VXSE53HypocenterArea `json:"area"`
	source string               `json:"source"`
}

type VXSE53HypocenterArea struct {
	Name                  string `json:"name"`
	Code                  string `json:"code"`
	CoordinateLatitude    string `json:"coordinate_latitude"`
	CoordinateLongitude   string `json:"coordinate_longitude"`
	CoordinateDepth       string `json:"coordinate_depth"`
	CoordinateDescription string `json:"coordinate_description"`
	DetailedName          string `json:"detailed_name"`
	DetailedCode          string `json:"detailed_code"`
	NameFromMark          string `json:"name_from_mark"`
	MarkCode              string `json:"mark_code"`
	Direction             string `json:"direction"`
	Distance              string `json:"distance"`
}

type VXSE53Intensity struct {
	Observation VXSE53Observation `json:"observation"`
}

type VXSE53Observation struct {
	MaxInt     string       `json:"max_int"`
	MaxIntCode string       `json:"max_int_code"`
	Pref       []VXSE53Pref `json:"pref"`
}

type VXSE53Pref struct {
	Name       string                `json:"name"`
	Code       string                `json:"code"`
	MaxInt     string                `json:"max_int"`
	MaxIntCode string                `json:"max_int_code"`
	Revise     string                `json:"revise"`
	Area       []VXSE53IntensityArea `json:area`
}

type VXSE53IntensityArea struct {
	Name       string       `json:"name"`
	Code       string       `json:"code"`
	MaxInt     string       `json:"max_int"`
	MaxIntCode string       `json:"max_int_code"`
	Revise     string       `json:"revise"`
	City       []VXSE53City `json:area`
}

type VXSE53City struct {
	Name             string                   `json:"name"`
	Code             string                   `json:"code"`
	MaxInt           string                   `json:"max_int"`
	MaxIntCode       string                   `json:"max_int_code"`
	Revise           string                   `json:"revise"`
	IntensityStation []VXSE53IntensityStation `json:"intensity_station"`
}

type VXSE53IntensityStation struct {
	Name       string `json:"name"`
	Code       string `json:"code"`
	MaxInt     string `json:"max_int"`
	MaxIntCode string `json:"max_int_code"`
	Revise     string `json:"revise"`
}

type VXSE53Comments struct {
	ForecastComment VXSE53ForecastComment `json:"forecast_comment"`
	VarComment      VXSE53VarComment      `json:"var_comment"`
	FreeFormComment string                `json:"free_form_comment"`
}

type VXSE53ForecastComment struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

type VXSE53VarComment struct {
	Text string `json:"text"`
	Code string `json:"code"`
}

// slack通知用メッセージを作る
func (data VXSE53ReceivedData) MakeNotifyMessage() string {
	var notify_message string

	notify_message = "発表日時：" + data.ReportDatetime.Format("2006-01-02 15:04") + "\n" +
	                 "発生時刻：" + data.Body.EarthQuake.OriginTime.Format("2006-01-02 15:04") + "\n" +
									 "震央地名：" + data.Body.EarthQuake.Hypocenter.Area.Name + "\n" +
									 "震源地マグニチュード文字列" + data.Body.EarthQuake.MagnitudeDescription + "\n" +
									 "最大震度：" + data.Body.Intensity.Observation.MaxInt + "\n" +
									 "各都道府県の最大震度は以下の通り：\n"

	for _, v := range data.Body.Intensity.Observation.Pref {
		notify_message += "     " + v.Name + "の最大震度：" + v.MaxInt + "\n"
	}

	return notify_message
}