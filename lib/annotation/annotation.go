package annotation

import (
	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/urfave/cli"
)

type ActionType int

const (
	LABEL_AS_POSITIVE ActionType = iota
	LABEL_AS_NEGATIVE
	HELP
	SKIP
	EXIT
)

func rune2ActionType(r rune) ActionType {
	switch r {
	case 'p':
		return LABEL_AS_POSITIVE
	case 'n':
		return LABEL_AS_NEGATIVE
	case 's':
		return SKIP
	case 'h':
		return HELP
	case 'e':
		return EXIT
	default:
		return HELP
	}
}

func NextExampleToBeAnnotated(m classifier.MIRAClassifier, examples model.Examples) *model.Example {
	unlabeledExamples := m.SortByScore(examples)
	if len(unlabeledExamples) == 0 {
		return nil
	}
	e := unlabeledExamples[0]
	if e == nil {
		return nil
	}
	return e
}

var ActionHelpDoc = `
p: Label this example as positive.
n: Label this example as negative.
s: Skip this example.
h: Show this help.
e: Exit.
`

var CommandAnnotate = cli.Command{
	Name:  "annotate",
	Usage: "Annotate URLs",
	Description: `
Annotate URLs using active learning.
`,
	Subcommands: []cli.Command{
		{
			Name:  "cli",
			Usage: "Annotate URLs using cli",
	