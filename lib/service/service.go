package service

import (
	"bufio"
	"io"
	"time"

	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/repository"
)

type GoActiveLearningApp interface {
	UpdateOrCreateExample(e *model.Example) error
	UpdateScore(e *model.Example) error
	InsertExampleFromScanner(scanner *bufio.Scanner) (*model.Examp