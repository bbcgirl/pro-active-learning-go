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

func TestPing(t *testing.