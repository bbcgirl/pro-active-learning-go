package repository

import (
	"time"

	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/model"
)

func (r *repository) UpdateOrCreateReferringTweets(e *model.Example) error {
	if e.Referring