
package service

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"
	"time"

	"math"
	"os"
	"strconv"
	"sync"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/fetcher"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/util"
)

func (app *goActiveLearningApp) UpdateOrCreateExample(e *model.Example) error {
	return app.repo.UpdateOrCreateExample(e)
}

func (app *goActiveLearningApp) UpdateScore(e *model.Example) error {
	return app.repo.UpdateScore(e)
}

func (app *goActiveLearningApp) InsertExampleFromScanner(scanner *bufio.Scanner) (*model.Example, error) {
	return app.repo.InsertExampleFromScanner(scanner)
}

func (app *goActiveLearningApp) InsertExamplesFromReader(reader io.Reader) error {
	return app.repo.InsertExamplesFromReader(reader)
}

func (app *goActiveLearningApp) SearchExamples() (model.Examples, error) {
	return app.repo.SearchExamples()
}

func (app *goActiveLearningApp) SearchRecentExamples(from time.Time, limit int) (model.Examples, error) {
	return app.repo.SearchRecentExamples(from, limit)
}

func (app *goActiveLearningApp) SearchRecentExamplesByHost(host string, from time.Time, limit int) (model.Examples, error) {
	return app.repo.SearchRecentExamplesByHost(host, from, limit)
}

func (app *goActiveLearningApp) SearchExamplesByLabel(label model.LabelType, limit int) (model.Examples, error) {
	return app.repo.SearchExamplesByLabel(label, limit)
}

func (app *goActiveLearningApp) SearchLabeledExamples(limit int) (model.Examples, error) {
	return app.repo.SearchLabeledExamples(limit)
}

func (app *goActiveLearningApp) SearchPositiveExamples(limit int) (model.Examples, error) {
	return app.repo.SearchPositiveExamples(limit)
}

func (app *goActiveLearningApp) SearchNegativeExamples(limit int) (model.Examples, error) {
	return app.repo.SearchNegativeExamples(limit)
}

func (app *goActiveLearningApp) SearchUnlabeledExamples(limit int) (model.Examples, error) {
	return app.repo.SearchUnlabeledExamples(limit)
}

func (app *goActiveLearningApp) SearchPositiveScoredExamples(limit int) (model.Examples, error) {
	return app.repo.SearchPositiveScoredExamples(limit)
}

func (app *goActiveLearningApp) FindExampleByUlr(url string) (*model.Example, error) {
	return app.repo.FindExampleByUlr(url)
}

func (app *goActiveLearningApp) FindExampleById(id int) (*model.Example, error) {
	return app.repo.FindExampleById(id)
}

func (app *goActiveLearningApp) SearchExamplesByUlrs(urls []string) (model.Examples, error) {
	return app.repo.SearchExamplesByUlrs(urls)
}

func (app *goActiveLearningApp) SearchExamplesByIds(ids []int) (model.Examples, error) {
	return app.repo.SearchExamplesByIds(ids)
}

func (app *goActiveLearningApp) SearchExamplesByKeywords(keywords []string, aggregator string, limit int) (model.Examples, error) {
	return app.repo.SearchExamplesByKeywords(keywords, aggregator, limit)
}

func (app *goActiveLearningApp) DeleteAllExamples() error {
	return app.repo.DeleteAllExamples()
}

func (app *goActiveLearningApp) CountPositiveExamples() (int, error) {
	return app.repo.CountPositiveExamples()
}

func (app *goActiveLearningApp) CountNegativeExamples() (int, error) {
	return app.repo.CountNegativeExamples()
}

func (app *goActiveLearningApp) CountUnlabeledExamples() (int, error) {
	return app.repo.CountUnlabeledExamples()
}

func (app *goActiveLearningApp) UpdateFeatureVector(e *model.Example) error {
	return app.repo.UpdateFeatureVector(e)
}

func (app *goActiveLearningApp) UpdateHatenaBookmark(e *model.Example) error {
	return app.repo.UpdateHatenaBookmark(e)
}

func (app *goActiveLearningApp) UpdateOrCreateReferringTweets(e *model.Example) error {
	return app.repo.UpdateOrCreateReferringTweets(e)
}

func (app *goActiveLearningApp) UpdateTweetLabel(exampleId int, idStr string, label model.LabelType) error {
	return app.repo.UpdateTweetLabel(exampleId, idStr, label)
}

func (app *goActiveLearningApp) SearchReferringTweets(limit int) (model.ReferringTweets, error) {
	return app.repo.SearchReferringTweets(limit)
}

func (app *goActiveLearningApp) SearchPositiveReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, limit int) (model.ReferringTweets, error) {
	return app.repo.SearchPositiveReferringTweets(scoreThreshold, tweetsLimitInSameExample, limit)
}

func (app *goActiveLearningApp) SearchNegativeReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, limit int) (model.ReferringTweets, error) {
	return app.repo.SearchNegativeReferringTweets(scoreThreshold, tweetsLimitInSameExample, limit)
}

func (app *goActiveLearningApp) SearchUnlabeledReferringTweets(scoreThreshold float64, tweetsLimitInSameExample int, limit int) (model.ReferringTweets, error) {
	return app.repo.SearchUnlabeledReferringTweets(scoreThreshold, tweetsLimitInSameExample, limit)
}

func (app *goActiveLearningApp) SearchRecentReferringTweetsWithHighScore(from time.Time, scoreThreshold float64, limit int) (model.ReferringTweets, error) {
	return app.repo.SearchRecentReferringTweetsWithHighScore(from, scoreThreshold, limit)
}

func hatenaBookmarkByExampleId(hatenaBookmarks []*model.HatenaBookmark) map[int]*model.HatenaBookmark {
	result := make(map[int]*model.HatenaBookmark)
	for _, hb := range hatenaBookmarks {
		result[hb.ExampleId] = hb
	}
	return result
}

func (app *goActiveLearningApp) AttachMetadataIncludingFeatureVector(examples model.Examples, bookmarkLimit int, tweetLimit int) error {
	// make sure that example id must be filled
	for _, e := range examples {
		if e.Id == 0 {
			tmp, err := app.FindExampleByUlr(e.Url)
			if err != nil {
				return err
			}
			e.Id = tmp.Id
		}
	}

	fvList, err := app.repo.SearchFeatureVector(examples)
	if err != nil {
		return err
	}

	for _, e := range examples {
		if fv, ok := fvList[e.Id]; ok {
			e.Fv = fv
		}
	}

	return app.AttachMetadata(examples, bookmarkLimit, tweetLimit)
}

func (app *goActiveLearningApp) AttachMetadata(examples model.Examples, bookmarkLimit int, tweetLimit int) error {
	hatenaBookmarks, err := app.repo.SearchHatenaBookmarks(examples, bookmarkLimit)
	if err != nil {
		return err
	}
	hbByid := hatenaBookmarkByExampleId(hatenaBookmarks)