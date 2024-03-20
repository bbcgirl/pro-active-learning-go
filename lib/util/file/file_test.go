package file

import (
	"fmt"
	"testing"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
)

func TestParseLine(t *testing.T) {
	line1 := "http://model.com\t1"
	e, err := ParseLine(line1)

	if err != nil {
		t.Error("cannot parse line1")
	}
	if e.Label != model.POSITIVE {
		t.Error("Label must be POSITIVE")
	}

	line2 := "http://model.com\t-1"
	e, err = ParseLine(line2)

	if err != nil {
		t.Error("cannot pa