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
	tweet := model.Tweet{}
	tweet.ScreenName = "syou6162"
	tweet.FullText = "Hello world @syou6162 @syou6163 #hashtag1 #hashtag2"
	tweet.FavoriteCount = 7
	tweet.RetweetCount = 7

	et := 