package service

import (
	"context"

	"github.com/b92c/ineed-engine/pkg/database"
)

type FreelanceSearchService struct {
	db database.DB
}

func NewFreelanceSearchService(db database.DB) *FreelanceSearchService {
	return &FreelanceSearchService{db: db}
}

func (s *FreelanceSearchService) GetHealthStatus(ctx context.Context) (map[string]string, error) {
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
