package repository

import (
	"time"

	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/model"
)

func (r *repository) UpdateOrCreateReferringTweets(e *model.Example) error {
	if e.ReferringTweets == nil || len((*e).ReferringTweets.Tweets) == 0 || (*e).ReferringTweets.Count == 0 {
		return nil
	}

	tmp, err := r.FindExampleByUlr(e.Url)
	if err != nil {
		return err
	}
	id := tmp.Id

	for _, t := range (*e).ReferringTweets.Tweets {
		t.ExampleId = id
		if _, err = r.db.NamedExec(`
INSERT INTO tweet
( example_id,  created_at,  id_str,  full_text,  favorite_count,  retweet_count,  lang,  screen_name,  name,  profile_image_url,  label,  score)
VALUES
(:example_id, :created_at, :id_str, :full_text, :favorite_count, :retweet_count, :lang, :screen_name, :name, :profile_image_url, :label, :score)
ON CONFLICT (example_id, id_str)
DO UPDATE SET
favorite_count = :favorite_count,  retweet_count = :retweet_count, label = :label
WHERE
EXCLUDED.label != 0 AND tweet.label != EXCLUDED.label
;`, t); err != nil {
			return err
		}
	}
	return nil
}

func (r *repository) UpdateTweetLabel(exampleId int, idStr string, label model.LabelType) error {
	if _, err := r.db.Exec(`UPDATE tweet SET label = $1 WHERE example_id = $2 AND id_str = $3;`, label, exampleId, idStr); err != nil {
		return err
	}
	return nil
}

type exampleIdWithTweetsCount struct {
	ExampleId   int `db:"example_id"`
	TweetsCount int `db:"tweets_count"`
}

func (r *repository) SearchReferringTweetsList(examples model.Examples, limitForEachExample int) (map[int]model.ReferringTweets, error) {
	referringTweetsByExampleId := make(map[int]model.ReferringTweets)

	exampleIds := make([]int, 0)
	for _, e := range examples {
		exampleIds = append(exampleIds, e.Id)
	}

	exampleIdsWithTweetsCount := make([]exampleIdWithTweetsCount, 0)
	tweetsCountByExampleQuery := `SELECT example_id, COUNT(*) AS tweets_count FROM tweet WHERE example_id = ANY($1) GROUP BY example_id ORDER BY tweets_count DESC;`
	err := r.db.Select(&exampleIdsWithTweetsCount, tweetsCountByExampleQuery, pq.Array(exampleIds))
	if err != nil {
		return referringTweetsByExampleId, err
	}
	tweetsCountByExampleId := make(map[int]int)
	for _, e := range exampleIdsWithTweetsCount {
		tweetsCountByExampleId[e.ExampleId] = e.TweetsCount
	}

	if limitForEachExample == 0 {
		for _, exampleId := range exampleIds {
			referringTweets := model.ReferringTweets{Count: 0, Tweets: make([]*model.Tweet, 0)}
			if cnt, ok := tweetsCountByExampleId[exampleId]; ok {
				referringTweets.Count = cnt
			}
			referringTweetsByExampleId[exampleId] = referringTweets
		}
		return referringTweetsByExampleId, nil
	}

	tweets := make([]*model.Tweet, 0)
	query := `SELECT * FROM tweet WHERE example_id = ANY($1) AND label != -1 AND score > -1.0 AND 