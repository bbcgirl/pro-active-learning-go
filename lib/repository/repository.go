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
	SearchRecentExamplesByHost(host string, from time.Time, limit int) (model.Examples, error)
	SearchExamplesByLabel(label model.LabelType, limit int) (model.Examples, error)
	SearchLabeledExamples(limit int) (model.Examples, error)
	SearchPositiveExamples(limit int) (model.Examples, error)
	SearchNegativeExamples(limit int) (model.Examples, error)
	SearchUnlabeledExamples(limit int) (model.Examples, error)
	SearchPositiveScoredExamples(limit int) (model.Examples, error)
	FindExampleByUlr(url string) (*model.Example, error)
	FindExampleById(id int) (*model.Example, error)
	SearchExamplesByUlrs(urls []string) (model.Examples, error)
	SearchExamplesByIds(ids []int) (model.Examples, error)
	SearchExamplesByKeywords(keywords []string, aggregator string, limit int) (model.Examples, error)
	DeleteAllExamples() error

	CountPositiveExamples() (int, error)
	CountNegativeExamples() (int, error)
	CountUnlabeledExamples() (int, error)

	IncErrorCount(e *model.Example) error
	GetErrorCount(e *model.Example) (int, error)

	UpdateFeatureVector(e *model.Example) error
	FindFeatureVector(e *model.Example) (feature.FeatureVector, error)
	SearchFeatureVector(examples model.Examples) (map[int]feature.FeatureVector, error)

	UpdateHatenaBookmark(e *model.Example) error
	SearchHatenaBookmarks(examples model.Examples, limitForEachExample int) ([]*model.HatenaBookmark, error)
	FindHatenaBookmark(e *model.Example, limit int) (*model.HatenaBookmark, error)

	UpdateOrCreateReferringTweets(e *model.Example) error
	UpdateTweetLabel(exampleId int, idStr string, label model.LabelType) error
	SearchReferringTweetsList(examples model.Examples, limit int) (map[int]model.ReferringTweets, error)
	SearchReferringTweets(limit int) (model.ReferringTweets, error)
	SearchPositiveReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, limit int) (model.ReferringTweets, error)
	SearchNegativeReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, l