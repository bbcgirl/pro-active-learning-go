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
	res, err := http.Get(fmt.Sprintf("htt