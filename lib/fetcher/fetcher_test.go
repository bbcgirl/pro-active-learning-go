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
	if err != nil {
		t.Error(err.Error())
	}

	if a.Title != "[2012.07805] Extracting Training Data from Large Language Models" {
		t.Error("Title must not be empty")
	}
	if a.Description == "" {
		t.Error("Description must not be empty")
	}
	if a.StatusCode != 200 {
		t.Error("StatusCode must be 200")
	}
}

func TestGetArticleNotFound(t *testing.T) {
	_, err := GetArticle("https://www.yasuhisay.info/entry/NOT_FOUND")
	if err == nil {
		t.Error("Error should occur")
	}
}

func TestGetArticleWithInvalidEncoding(t *testing.T) {
	url := "http://www.atmarkit.co.jp/ait/articles/1702/20/news021.html"
	_, err := GetArticle(url)
	if err == nil {
		t.Error(fmt.Sprintf("Error must occur for this url: %s", url))
	}
}

func TestRemoveUtmParams(t *testing.T) {
	before := "https://techplay.jp/event/698349?utm_source=event_698349"
	after, err := removeUtmParams(before)
	if err != nil {
		t.Error(fmt.Sprintf("Error must occur for this url: %s", before))
	}
	expected := "https://techplay.jp/event/698349"
	if expected != after {
		t.Errorf("url should be %s, but %s", expected, after)
	}
	a, err := GetArticle(before)
	if expected != a.Url {
		t.Errorf("url should be %s, but %s", expected, a.Url)
	}
}

func TestFavicon(t *testing.T) {
	url := "https://ww