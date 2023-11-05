package repository

import (
	"fmt"
	"io"
	"time"

	"github.com/jmoiron/sqlx"

	"bufio"

	_ "github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/feature"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/util"
)

type Repository interface {
	UpdateOrCreateExample(e *model.Example) error
	UpdateScore(e *model.Example) error
	InsertExampleFromScanner(scanner *bufio.Scanner) (*model.Example, error)
	InsertExamplesFromReader(reader io.Reader) error
	SearchExamples() (model.Examples, error)
	SearchRecentExamples(from time.Time, limit int) (model.Examples, error)
	SearchRecentExamplesByHost(host string, from time.Time, limit int) (model.Examples, erro