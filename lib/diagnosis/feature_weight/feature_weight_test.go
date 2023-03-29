package featureweight_test

import (
	"testing"

	"github.com/syou6162/go-active-learning/lib/command"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou6162/go-active-learning/lib/util/file"
	"github.com/urfave/cli"
)

func TestDoListFeatureWeight(t *testing.T) {
	inputFilename := "../..