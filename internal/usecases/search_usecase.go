package usecases

import (
	"context"
	"log"

	"github.com/b92c/ineed-engine/pkg/database"
)

type SearchUsecase struct {
	db database.DB
}

func NewSearchUsecase(db database.DB) *SearchUsecase {
	return &SearchUsecase{db: db}
}

func (s *SearchUsecase) Find(ctx context.Context) ([]string, error) {
	r, err := s.db.GetAll()
	if err != nil {
		log.Fatal("failed to get result from database: ", err)
		return nil, err
	}
	return r, nil
}
