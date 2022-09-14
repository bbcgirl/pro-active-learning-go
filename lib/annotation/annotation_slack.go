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
	filterStatusCodeOk := c.Bool("filter-status