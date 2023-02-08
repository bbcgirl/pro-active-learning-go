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
	GetFeatureVector() feature.FeatureVector
	GetLabel() model.LabelType
}

type LearningInstances []LearningInstance

var errNoTrainingInstances = errors.New("Empty training set")
var errNoDevelopmentInstances = errors.New("Empty development set")
var errNoMIRAModelLearned = errors.New("Fail to learn MIRA models")
var errModelEvaluationFailure = errors.New("Failed to evaluate best MIRA")
var errTrainingInstancesAllPositive = errors.New("Labels of training instances are all positive")
var errTrainingInstancesAllNegative = errors.New("Labels of training instances are all negative")
var errDevelopmentInstancesAllPositive = errors.New("Labels of development instances are all positive")
var errDevelopmentInstancesAllNegative = errors.New("Labels of development instances are all negative")

func newMIRAClassifier(modelType ModelType, c float64) *MIRAClassifier {
	return &MIRAClassifier{
		ModelType: modelType,
		Weight:    make(map[string]float64),
		C:         c,
		Accuracy:  0.0,
		Precision: 0.0,
		Recall:    0.0,
		Fvalue:    0.0,
	}
}

func filterLabeledInstances(instances LearningInstances) LearningInstances {
	var result LearningInstances
	for _, i := range instances {
		if i.GetLabel() != 0 {
			result = append(result, i)
		}
	}
	return result
}

func shuffle(instances LearningInstances) {
	n := len(instances)
	for i := n - 1; i >= 0; i-- {
		j := rand.Intn(i + 1)
		instances[i], instances[j] = instances[j], instances[i]
	}
}

func splitTrainAndDev(instances LearningInstances) (train LearningInstances, dev LearningInstances) {
	shuffle(instances)
	n := int(0.8 * float64(len(instances)))
	return instances[0:n], instances[n:]
}

func NewMIRAClassifier(modelType ModelType, instances LearningInstances, 