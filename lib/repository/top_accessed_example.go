package repository

import (
	"github.com/lib/pq"
)

func (r *repository) UpdateTopAccessedExampleIds(exampleIds []int) error {
	if _, err := r.db.Exec(`DELETE FROM top_accessed_example;`); err != nil {
		return err
	}
	if _, err := r.db.Exec(`INSERT INTO top_accessed_example (example_id) VALUES (unnest(cast($1 AS INT[])));`, pq.Array(exampleIds)); err != nil {
		return err
	}
	return nil
}

func (r *rep