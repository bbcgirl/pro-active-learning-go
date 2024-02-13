package service

import (
	"bufio"
	"io"
	"time"

	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/repository"
)

type GoActiveLearningApp interface {
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

	InsertMIRAModel(m classifier.MIRAClassifier) error
	FindLatestMIRAModel(modelType classifier.ModelType) (*classifier.MIRAClassifier, error)

	UpdateFeatureVector(e *model.Example) error
	UpdateHatenaBookmark(e *model.Example) error
	UpdateOrCreateReferringTweets(e *model.Example) error
	UpdateTweetLabel(exampleId int, idStr string, label model.LabelType) error
	SearchReferringTweets(limit int) (model.ReferringTweets, error)
	SearchPositiveReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, limit int) (model.ReferringTweets, error)
	SearchNegativeReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, limit int) (model.ReferringTweets, error)
	SearchUnlabeledReferringTweets(scoreThreshold float64, tweetsLimitInSa