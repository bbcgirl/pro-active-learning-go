package repository

import (
	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/model"
)

func (r *repository) UpdateRecommendation(rec model.Recommendation) error {
	if _, err := r.db.Exec(`DELETE FROM recommendation WHERE list_type = $1;`, rec.RecommendationListType); err != nil {
		ret