package model

import (
	"database/sql/driver"
	"encoding/json"
	"strings"
	"time"
)

type Tags []string

type HatenaBookmarkTime struct {
	*time.Time
}

// ref: https://dev.classmethod.jp/go/struct-json/
func (hbt *HatenaBookmarkTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse("\"2006/01/02 15:04\"", string(data))
	*hbt = HatenaBookmarkTime{&t}
	return err
}

func (hbt HatenaBookmarkTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(hbt.Format("2006/01/02 15:04"))
}

// ref: https://qiita.com/roothybrid7/items/52623bedb45ff0c26a8a
func (hbt *HatenaBookmarkTime) Scan(value interface{}) error {
	v := value.(time.Time)
	hbt.Time = &v
	return nil
}

func (hbt Hate