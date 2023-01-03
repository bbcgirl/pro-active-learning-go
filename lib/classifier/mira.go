package classifier

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"sync"

	"github.com/pkg/errors"
	"github.com/syou6162/go-active-learning/lib/evaluation"
	"github.com/syou6162/go-active-learning/lib/feature"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/util"
)

type ModelType int

const (
	EXAMPLE ModelType = 0
	TWITTER ModelType = 1
)

type MIRAClassifier struct {
	ModelType ModelType          `json:"ModelType"`
	Weight    map[string]float64 `json:"Weight"`
	C         float64            `json:"C"`
	Accuracy  float64            `json:"Accuracy"`
	Precision float64            `json:"Precision"`
	Recall    float64            `json:"Recall"`
	Fvalue    float64            `json:"Fvalue"`
}

type LearningInstance interface {
	GetFeatureVector() feature.FeatureVec