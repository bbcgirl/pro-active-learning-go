package tweet_feature

import (
	"reflect"
	"testing"

	"github.com/syou6162/go-active-learning/lib/feature"
	"github.com/syou6162/go-active-learning/lib/model"
)

func TestExtractHostFeature(t *testing.T) {
	e := model.Example{}
	e.Title = "Hello world"
	tw