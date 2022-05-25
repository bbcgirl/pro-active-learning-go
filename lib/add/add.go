package add

import (
	"fmt"
	"log"
	"time"

	"os"

	mkr "github.com/mackerelio/mackerel-client-go"
	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/hatena_bookmark"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou6162/go-active-learning/lib/util"
	"github.com/syou6162/go-active-learning/lib/util/file"
	"github.com/urfave/cli"
)

func doAdd(c *cli.Context) error {
	inputFilename := c.String("input-filename")

	if inputFilename == "" {
		_ = cli.ShowCommandHelp(c, "add")
		return cli.NewExitError("`input-filename` is a required field.", 1)
	}

	app, err := service.NewDefaultApp()
	if err != nil {
		return err
	}
	defer app.Close()

	examples, err := file.ReadExamples(inputFilename)
	if err != nil {
		return err
	}

	if err := app.AttachMetadata(examples, 0, 0); err != nil {
		return err
	}

	examples = util.FilterStatusCodeNotOkExamples(examples)
	app.Fetch(examples)
	examples = util.FilterStatusCodeOkExamples(examples)

	m, err := app.FindLatestMIRAModel(classifier.EXAMPLE)
	skipPredictScore := false
	if err != nil {
		log.Println(fmt.Sprintf("Error to load model %s", err.Error()))
		skipPredictScore = true
	}

	for _, e := range examples {
		if !skipPredictScore {
			e.Score = m.PredictScore(e.Fv)
		}
		if e.CreatedAt.Before(time.Date(2000, 01, 01, 0, 0, 0, 0, time.Local)) {
			log.Println(fmt.Sprintf("Skipin too old example: %s", e.Url))
			continue
		}
		if err = app.UpdateOrCreateExample(e); err != nil {
			log.Println(fmt.Sprintf("Error occured proccessing %s %s", e.Url, err.Error()))
			continue
		}
		if err = app.UpdateFeatureVector(e); err != nil {
			log.Println(fmt.Sprintf("Error occured proccessing %s feature vector %s", e.Url, err.Error()))
			continue
		}
		if bookmark, err := hatena_bookmark.Get