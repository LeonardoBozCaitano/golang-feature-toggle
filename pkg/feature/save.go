package feature

import (
	"context"
)

func (s *Service) Save(ctx context.Context, data Feature) (int, error) {
	var id int
	row, err := s.db.NamedQuery("INSERT INTO feature(name, responsible, active) VALUES (:name, :responsible, true) RETURNING id", data)
	if err != nil {
		return 0, err
	}
	if row.Next() {
		row.Scan(&id)
	}
	return id, nil
}
