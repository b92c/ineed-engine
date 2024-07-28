package usecases

import (
	"context"

	"github.com/b92c/ineed-engine/pkg/database"
)

type HealthCheckUsecase struct {
	db database.DB
}

func NewHealthCheckUsecase(db database.DB) *HealthCheckUsecase {
	return &HealthCheckUsecase{db: db}
}

func (s *HealthCheckUsecase) GetHealthStatus(ctx context.Context) (map[string]string, error) {
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
