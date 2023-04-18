package example_feature

import (
	"fmt"
	"testing"
)

func TestIsJapanese(t *testing.T) {
	text := "ほげ"
	if !isJapanese(text) {
		t.Error(fmt.Printf("%s should be Japanese", text))
	}
	text = "文献紹介 / Youtube"
	if !isJapanese(text) {
		t.Error(fmt.Printf("%s should be Japanese", text))
	}
	text = "This is a pen."
	if isJapanese(text) {
		t.Error(fmt.Printf("%s should be not Japanese", text))
	}
}

func TestJapaneseNounFeatures(t *testing.T) {
	text := "日本語のテストです"
	fv := ExtractJpnNounFeaturesWithoutPrefix(text)
	if len(fv) != 2 {
		t.Error(fmt.Printf("Size of feature vector for %s 