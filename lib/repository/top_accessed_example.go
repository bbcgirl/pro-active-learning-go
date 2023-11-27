package repository

import (
	"github.com/lib/pq"
)

func (r *repository) UpdateTopAccessedExampleIds(exampleIds []int) error {
	if _, err := r.db.Exec(`DELETE FROM top_accessed_example;`); err != nil {
		return err
	}
	if _, err :