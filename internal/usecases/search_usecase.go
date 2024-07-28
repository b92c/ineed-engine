package usecases

import (
	"context"

	"github.com/b92c/ineed-engine/pkg/database"
)

type SearchUsecase struct {
	db database.DB
}

func NewSearchUsecase(db database.DB) *SearchUsecase {
	return &SearchUsecase{db: db}
}

func (s *SearchUsecase) GetHealthStatus(ctx context.Context) (map[string]string, error) {
	stats := make(map[string]string)
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = err.Error()
		return stats, err
	}
	stats["status"] = "up"
	return stats, nil
}
