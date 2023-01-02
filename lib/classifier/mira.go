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
	ModelType Mode