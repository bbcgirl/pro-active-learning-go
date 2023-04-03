package featureweight_test

import (
	"testing"

	"github.com/syou6162/go-active-learning/lib/command"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou6162/go-active-learning/lib/util/file"
	"github.com/urfave/cli"
)

func TestDoListFeatureWeight(t *testing.T) {
	inputFilename := "../../../tech_input_example.txt"
	train, err := file.ReadExamples(inputFilename)
	if err != nil {
		t.Error(err)
	}

	a, err := service.NewDefaultApp()
	if err != nil {
		t.Error(err)
	}
	defer a.Close()

	if err = a.DeleteAllExamples(); err != nil {
		t.Error(err)
	}

	for _, example := range train {
		if err = a.UpdateOrCreateExample(example); err != nil {
			t.Error(err)
		}
	}

	app :=