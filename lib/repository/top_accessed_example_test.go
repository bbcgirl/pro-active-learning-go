package repository_test

import (
	"testing"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/repository"
)

func TestUpdateTopAccessedExampleIds(t *testing.T) {
	repo, err := repository.New()
	if err != nil {
