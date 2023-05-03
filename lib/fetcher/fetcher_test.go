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

	if a.Title == "" {
		t.Error("Title must not be empty")
	}
	if a.Description == "" {
		t.Error("Description must not be empty")
	}
	if a.OgType != "article" {
		t.Error("OgType must be article")
	}
	if a.StatusCode != 200 {
		t.Error("StatusCode must be 200")
	}
}

func TestGetArticleARXIV(t *testing.T) {
	a, err := GetArticle("https://arxiv.org/abs/2012.07805")
	i