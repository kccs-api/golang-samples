package endpoint

import (
	"fmt"
	"net/http"
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
    fmt.Printf("%+v\n", pr)

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
	fmt.Println(string(json_str))

	// データ部分のJSON文字列を構造体に変換する
	var rd structs.ReceivedData
	json.Unmarshal([]byte(json_str), &rd)

	// 変換結果を出力する
	fmt.Printf("%+v\n", rd)

	// 200OKのレスポンスを返す
	fmt.Fprint(w, "OK")
}