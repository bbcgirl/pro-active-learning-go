package featureweight

import (
	"fmt"
	"sort"

	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou6162/go-active-learning/lib/util"
	"github.com/syou6162/go-active-learning/lib/util/converter"
	"github.com/urfave/cli"
)

type Feature struct {
	Key    string
	