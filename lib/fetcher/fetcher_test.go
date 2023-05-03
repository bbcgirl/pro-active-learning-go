package fetcher

import (
	"fmt"
	"testing"
)

func TestGetArticle(t *testing.T) {
	a, err := GetArticle("https://www.yasuhisay.info/entry/20090516/1242480413")
	if err != nil {
		t.Error(err.Error())
	}

	if a.Title =