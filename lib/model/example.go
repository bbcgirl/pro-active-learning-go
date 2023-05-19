
package model

import (
	"math"
	"strings"
	"time"

	"github.com/syou6162/go-active-learning/lib/feature"
)

type Example struct {
	Id              int       `db:"id"`
	Label           LabelType `json:"Label" db:"label"`
	Fv              feature.FeatureVector
	Url             string           `json:"Url" db:"url"`
	FinalUrl        string           `json:"FinalUrl" db:"final_url"`
	Title           string           `json:"Title" db:"title"`
	Description     string           `json:"Description" db:"description"`
	OgDescription   string           `json:"OgDescription" db:"og_description"`
	OgType          string           `json:"OgType" db:"og_type"`
	OgImage         string           `json:"OgImage" db:"og_image"`
	Body            string           `json:"Body" db:"body"`
	Score           float64          `db:"score"`
	IsNew           bool             `db:"is_new"`
	StatusCode      int              `json:"StatusCode" db:"status_code"`
	Favicon         string           `json:"Favicon" db:"favicon"`
	ErrorCount      int              `json:"ErrorCount" db:"error_count"`
	CreatedAt       time.Time        `json:"CreatedAt" db:"created_at"`
	UpdatedAt       time.Time        `json:"UpdatedAt" db:"updated_at"`
	ReferringTweets *ReferringTweets `json:"ReferringTweets"`
	HatenaBookmark  *HatenaBookmark  `json:"HatenaBookmark"`
}

type Examples []*Example

func (example *Example) GetLabel() LabelType {
	return example.Label
}

func (example *Example) GetFeatureVector() feature.FeatureVector {
	return example.Fv
}

func (example *Example) Annotate(label LabelType) {
	example.Label = label
}
