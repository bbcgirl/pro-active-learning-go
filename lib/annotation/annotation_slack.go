package annotation

import (
	"fmt"
	"os"

	"github.com/nlopes/slack"
	"github.com/pkg/errors"
	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou6162/go-active-learning/lib/util"
	"github.com/syou6162/go-active-learning/lib/util/converter"
	"github.com/urfave/cli"
)

func doAnnotateWithSlack(c *cli.Context) error {
	channelID := c.String("channel")
	filterStatusCodeOk := c.Bool("filter-status-code-ok")

	if channelID == "" {
		_ = cli.ShowCommandHelp(c, "slack")
		return cli.NewExitError("`channel` is a required field.", 1)
	}

	api := slack.New(os.Getenv("SLACK_TOKEN"))
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	app, err := service.NewDefaultApp()
	if err != nil {
		return err
	}
	defer app.Close()

	examples, err := app.SearchExamples()
	if err != nil {
		return err
	}

	stat := example.GetStat(examples)
	msg := rtm.NewOutgoingMessage(fmt.Sprintf("Positive:%d, Negative:%d, Unlabeled:%d", stat["positive"], stat["negative"], stat["unlabeled"]), channelID)
	rtm.SendMessage(msg)

	app.Fetch(examples)
	for _, e := range examples {
		app.UpdateFeatureVector(e)
	}
	if filterStatusCodeOk {
		examples = util.FilterStatusCodeOkExamples(examples)
	}

	m, err := classifier.NewMIRAClassifierByCrossValidation(classifier.EXAMPLE, converter.ConvertExamplesToLearningInstances(examples))
	if err != nil {
		return err
	}
	e := NextExampleToBeAnnotated(*m, examples)
	if e == nil {
		return errors.New("No e to annotate")
	}

	rtm.SendMessage(rtm.NewOutgoingMessage("Ready to annotate!", channelID))
	showExample(rtm, *m, e, channelID)
	prevTimestamp := ""

annotationLoop:
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.AckMessage:
				prevTimestamp = ev.Timestamp
			case *slack.MessageEvent:
				if ev.Channel != channelID {
					break
				}
				text := ev.Text
				if len(text) > 1 || len(text