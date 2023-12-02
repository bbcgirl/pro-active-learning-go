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
favorite_count = :favorite_count,  retweet_count = :retweet_co