package structs

import (
	"time"
	"regexp"
	"encoding/json"
)

// time.Time型をそのまま使うと空の文字列をパースする際にエラーとなるため、独自に構造体を定義する
type AllowEmptyTime struct {
	*time.Time
}

// 文字列 -> 時刻に変換する際、空文字の場合は 0001-01-01 00:00:00 +0000 UTC として扱う
func (aet *AllowEmptyTime) UnmarshalJSON(data []byte) error {
	r := regexp.MustCompile(`""`)
	if r.MatchString(string(data)) {
		*aet = AllowEmptyTime{&time.Time{}}
		return nil
	}

	t, err := time.Parse("\""+time.RFC3339+"\"", string(data))
	*aet = AllowEmptyTime{&t}

	return err
}

func (aet *AllowEmptyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(aet.Format(time.RFC3339))
}
