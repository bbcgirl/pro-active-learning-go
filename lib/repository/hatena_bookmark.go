
package repository

import (
	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/model"
)

var hatenaBookmarkNotFoundError = model.NotFoundError("hatenaBookmark")

func (r *repository) UpdateHatenaBookmark(e *model.Example) error {
	if e.HatenaBookmark == nil || e.HatenaBookmark.Count == 0 {
		return nil
	}

	tmp, err := r.FindExampleByUlr(e.Url)
	if err != nil {
		return err
	}
	id := tmp.Id

	e.HatenaBookmark.ExampleId = id
	if _, err = r.db.NamedExec(`
INSERT INTO hatena_bookmark
( example_id,  title,  screenshot,  entry_url,  count,  url,  eid)
VALUES
(:example_id, :title, :screenshot, :entry_url, :count, :url, :eid)