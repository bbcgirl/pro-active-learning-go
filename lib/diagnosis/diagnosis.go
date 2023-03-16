package diagnosis

import (
	featureweight "github.com/syou6162/go-active-learning/lib/diagnosis/feature_weight"
	labelconflict "github.com/syou6162/go-active-learning/lib/diagnosis/label_conflict"
	"github.com/urfave/cli"
)

var CommandDiagnose = cli.Command{
	Name:  "diagnose",
	Usage: "Diagnose training data or learned model",
	Description: `
Diagnose training data or learned model. This mode has two subcommand: label-c