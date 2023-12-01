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