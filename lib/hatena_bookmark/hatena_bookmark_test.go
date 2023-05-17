package hatena_bookmark

import (
	"testing"
)

func TestGetHatenaBookmark(t *testing.T) {
	bookmarks, err := GetHatenaBookmark(