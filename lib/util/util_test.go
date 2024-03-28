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

	examples := FilterLabeledExamples(model.Examples{e1, e2, e3})
	if len(examples) != 2 {
		t.Error("Number of labeled examples should be 2")
	}
}

func TestFilterUnlabeledExamples(t *testing.T) {
	e1 := example.NewExample("https://b.hatena.ne.jp", model.POSITIVE)
	e2 := example.NewExample("https://www.yasuhisay.info", model.NEGATIVE)
	e3 := example.NewExample("http://google.com", model.UNLABELED)
	e3.Title = "Google"

	examples := FilterUnlabeledExamples(model.Examples{e1, e2, e3})
	if len(examples) != 1 {
		t.Error("Number of unlabeled examples should be 1")
	}
}

func TestFilterStatusCodeOkExamples(t *testing.T) {
	e1 := example.NewExample("https://b.hatena.ne.jp", model.POSITIVE)
	e1.StatusCode = 200
	e2 := example.NewExample("https://www.yasuhisay.info", model.NEGATIVE)
	e2.StatusCode = 404
	e3 := example.NewExample("http://google.com", model.UNLABELED)
	e3.StatusCode = 304

	examples := FilterStatusCodeOkExamples(model.Examples{e1, e2,