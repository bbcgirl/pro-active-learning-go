package repository_test

import (
	"testing"
	"time"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/repository"
)

func TestUpdateHatenaBookmark(t *testing.T) {
	repo, err := repository.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	defer repo.Close()

	if err = repo.DeleteAllExamples(); err != nil {
		t.Error(err)
	}

	e := example.NewExample("http://hoge.com", model.UNLABELED)
	err = repo.UpdateOrCreateExample(e)
	if err != nil {
		t.Error(err)
	}
	now := time.Now()
	b1 := model.Bookmark{
		User:      "syou6162",
		Comment:   "面白いサイトですね",
		Timestamp: model.HatenaBookmarkTime{Time: &now},
		Tags:      model.Tags{"hack"},
	}
	hb := model.HatenaBookmark{
		ExampleId: e.Id,
		Title:     "hoge",
		Count:     10,
		Bookmarks: []*model.Bookmark{&b1},
	}
	e.