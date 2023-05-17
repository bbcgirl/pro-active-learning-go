package hatena_bookmark

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/syou6162/go-active-learning/lib/model"
)

func GetHatenaBookmark(url string) (*model.HatenaBookmark, error) {
	// ref: http://developer.hatena.ne.jp/ja/documents/bookmark/apis/getinfo
	res, err := http.Get(fmt.Sprintf("https://b.hatena.ne.jp/entry/jsonlite/?url=%s", url))
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: %d", res.StatusCode)
	}

	defe