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
		t.Err