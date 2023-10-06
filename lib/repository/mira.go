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
	query := `INSERT INTO model (model_type, model, c, accuracy, precision, recall, fvalue) VALUES ($1, $2, $3, $4, $5, $6, $7);`
	if _, err := r.db.Exec(query, m.ModelType, string(byt