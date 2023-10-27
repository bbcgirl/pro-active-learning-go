package repository

import (
	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/model"
)

func (r *repository) UpdateRelatedExamples(related model.RelatedExamples) error {
	if _, err := r.db.Exec(`DELETE FROM related_example WHERE example_id = $1;`, related.E