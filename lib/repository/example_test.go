package repository_test

import (
	"log"
	"os"
	"testing"
	"time"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/feature"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/repository"
)

func TestMain(m *testing.M) {
	repo, err := repository.New()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer repo.Close()

	ret := m.Run()
	os.Exit(ret)
}

func TestPing(t *testing.T) {
	repo, err := repository.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	defer repo.Close()

	if err := repo.Ping(); err != nil {
		t.Errorf(err.Error())
	}
}

func TestInsertExamplesFromReader(t *testing.T) {
	repo, err := repository.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	defer repo.Close()

	if err = repo.DeleteAllExamples(); err != nil {
		t.Error(err)
	}

	fp, err := os.Open("../../tech_input_example.txt")
	defer fp.Close()
	if err != nil {
		t.Error(err)
	}
	repo.InsertExamplesFromReader(fp)

	examples, err := repo.SearchExamples()
	if err != nil {
		t.Error(err)
	}
	if len(examples) == 0 {
		t.Errorf("len(examples) > 0, but %d", len(examples))
	}
}

func TestInsertOrUpdateExample(t *testing.T) {
	repo, err := repository.New()
	if err != nil {
		t.Errorf(err.Error())
	}
	defer repo.Close()

	if err = repo.DeleteAllExamples(); err != nil {
		t.Error(err)
	}

	err = repo.UpdateOrCreateExample(example.NewExample("http://hoge.com", model.UNLABELED))
	if err != nil {
		t.Error(err)
	}

	examples, err := repo.SearchExamples()
	if err != nil {
		t.Error(err)
	}
	if len(examples) != 1 {
		t.Errorf("len(examples) == %d, want 1", len(examples))
	}
	if examples[0].Label != model.UNLABELED {
		t.Errorf("label == %d, want 0", examples[0].Label)
	}
	if examples[0].Id == 0 {
		t.Error("id must not be 0")
	}

	// same url
	err = repo.UpdateOrCreateExample(example.NewExample("http://hoge.com", model.NEGATIVE))
	if err != nil {
		t.Error(err)
	}

	examples, err = repo.SearchExamples()
	if err != nil {
		t.Error(err)
	}
	if len(examples) != 1 {
		t.Errorf("len(examples) == %d, want 1", len(examples))
	}
	if examples[0].Label != model.NEGATIVE {
		t.Errorf("label == %d, want -1", examples[0].Label)
	}

	// same url but different label
	err = repo.UpdateOrCreateExample(example.NewExample("http://hoge.com", model.POSITIVE))
	if err != nil {
		t.Error(err)
	}

	examples, err = repo.SearchExamples()
	if err != nil {
		t.Error(err)
	}
	if len(examples) != 1 {
		t.Errorf("len(examples) == %d, want 1", len(examples))
	}
	if examples[0].Label != model.POSITIVE {
		t.Errorf("label == %d, want 1", examples[0].Label)
	}

	// cannot update to unlabeled
	err = repo.UpdateOrCreateExample(examp