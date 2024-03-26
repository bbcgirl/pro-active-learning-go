package util

import (
	"testing"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
)

func TestFilterLabeledExamples(t *testing.T) {
	e1 := example.NewExample("https://b.hatena.ne.jp", model.POSITIVE)
	e2 := example.NewExample("https://www.yasuhisay.info", model.NEGATIVE)
	e3 := example.NewExample("http://google.com", model.UNLABELED)

	examples := FilterLabeledExamples(m