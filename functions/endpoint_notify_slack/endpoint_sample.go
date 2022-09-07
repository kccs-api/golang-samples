package endpoint

import (
	"os"
	"fmt"
	"bytes"
	"net/http"
	"io/ioutil"
	"encoding/json"
	"encoding/base64"
	"kccs.co.jp/endpoint/structs"
)

// リクエスト時この関数が実行される
func EndpointSample(w http.ResponseWriter, r *http.Request) {
	// Pubsubから配信されるデータの構造体を定義する
	var pr structs.PushRequest

	// 受信したJSON文字列を構造体に変換する
	// JSONの構成については右記参照: https://cloud.google.com/pubsub/docs/push#receiving_messages
	err := json.NewDecoder(r.Body).Decode(&pr)
	fmt.Printf("received_data: %+v\n", pr)

	// 変換に失敗した場合はBadRequestを返す
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}

	// 構造体のデータ部分をBase64デコードする
	json_str, err := base64.StdEncoding.DecodeString(pr.Message.Data)

	// 変換に失敗した場合はBadRequestを返す
	if err != nil {
		http.Error(w, fmt.Sprintf("Could not decode body: %v", err), http.StatusBadRequest)
		return
	}
	fmt.Println("json_str: "+string(json_str))

	// data_type_codeを取得する
	var rd_tmp map[string]interface{}
	json.Unmarshal(json_str, &rd_tmp)
	data_type_code := rd_tmp["data_type_code"].(string)

	var result int
	var message string

	// data_type_codeごとに処理を分岐する
	switch data_type_code {
	// 気象警報
	case "VXWW50":
		// var rd structs.VXWW50ReceivedData
		result = 200
		message = "OK"
		err = nil
	// 震度データ
	case "VXSE53":
		var rd structs.VXSE53ReceivedData
		json.Unmarshal(json_str, &rd)

		notify_message := rd.MakeNotifyMessage()

		result, message, err =  notifySlack(notify_message)
	}

	// functionのレスポンスを返す
	if err != nil {
		w.WriteHeader(500)
		fmt.Fprint(w, "Internal Server Error")
	} else {
		w.WriteHeader(result)
		fmt.Fprint(w, message)
	}
}

// slackに通知する
func notifySlack(message string) (int, string, error) {
	// リクエストボディを作る
	username := os.Getenv("USERNAME")
	icon_emoji := os.Getenv("ICON_EMOJI")
	message_json := map[string]string{"username": username, "icon_emoji": icon_emoji, "text": message}
	marshaled_json, err := json.Marshal(message_json)
	fmt.Println("marshaled_json: "+string(marshaled_json))
	converted_io := bytes.NewReader(marshaled_json)

	// リクエスト実行
	url := os.Getenv("SLACK_URL")
	req, err := http.NewRequest("POST", url, converted_io)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp,err := client.Do(req)

	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	if err != nil {
		return 0, "", err
	} else {
		return resp.StatusCode, string(body), err
	}
}