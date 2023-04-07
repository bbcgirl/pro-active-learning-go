package labelconflict

import (
	"fmt"
	"os"
	"sort"
	"strconv"

	"encoding/csv"

	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou6162/go-active-learning/lib/util"
	"github.com/syou6162/go-active-learning/lib/util/converter"
	"github.com/urfave/cli"
)

func DoLabelConflict(c *cli.Context) error {
	filterStatusCodeOk := c.Bool("filter-status-code-ok")

	app, err := service.NewDefaultApp()
	if err != nil {
		return err
	}
	defer app.Close()

	examples, err := app.SearchExamples()
	if err != nil {
		return err
	}
	app.Fetch(examples)
	for _, e := range examples {
		app.UpdateFeatureVector(e)
	}
	training := util.FilterLabeledExamples(examples)

	if filterStatusCodeOk {
		training = util.FilterStatusCodeOkExamples(training)
	}

	m, err := classifier.NewMIRAClassifierByCrossValidation(classifier.EXAMPLE, converter.ConvertExamplesToLearningInstances(training))
	if err != nil {
		return err
	}

	wrongExamples := model.Examples{}
	correctExamples := model