package annotation

import (
	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/urfave/cli"
)

type ActionType int

const (
	LABEL_AS_POSITIVE ActionTyp