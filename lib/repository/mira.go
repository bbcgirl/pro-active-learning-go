package repository

import (
	"encoding/json"

	"github.com/syou6162/go-active-learning/lib/classifier"
)

func (r *repository) InsertMIRAModel(m classifier.MIRAClassifier) error {
	bytes, err := json.Marshal(m)
	if err != nil {
		return err
	}
	q