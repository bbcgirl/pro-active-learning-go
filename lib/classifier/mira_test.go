package classifier

import (
	"testing"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
)

func TestPredictScore(t *testing.T) {
	e1 := example.NewExample("https://b.hatena.