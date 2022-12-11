package feature

import "context"

func (s *Service) GetById(ctx context.Context, id int) (Feature, error) {
	var item Feature
	err := s.db.Get(&item, "SELECT * FROM feature WHERE id = $1", id)
	return item, err
}
