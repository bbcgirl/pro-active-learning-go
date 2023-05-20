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

/