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
	Weight float64
}

type FeatureList []Feature

func (p FeatureList) Len() int           { return len(p) }
func (p FeatureList) Less(i, j int) bool { return p[i].Weight < p[j].Weight }
func (p Fe