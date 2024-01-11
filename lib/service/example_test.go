package service_test

import (
	"testing"

	"github.com/syou6162/go-active-learning/lib/example"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/service"
)

func findExampleByurl(examples model.Examples, url string) *model.Example {
	for _, e := range examples {
		if e.Url == url {
			return e
		}
	}
	return nil
}

func TestAttach