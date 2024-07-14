package service

import (
	"context"

	"github.com/b92c/ineed-engine/pkg/database"
)

type HealthCheckService struct {
	db database.DB
}

func NewHealthCheckService(db database.DB) *HealthCheckService {
	return &HealthCheckService{db: db}
}

func (s *HealthCheckService) GetHealthStatus(ctx context.Context) (map[string]string, error) {
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
