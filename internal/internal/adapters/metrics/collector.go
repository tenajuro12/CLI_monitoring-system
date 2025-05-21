package metrics

import (
	"context"
	"monitoring/internal/domain"
)

type CPUCollector interface {
	Collect(ctx context.Context) (domain.CPUMetrics, error)
}

type MemoryCollector interface {
	Collect(ctx context.Context) (domain.MemoryMetrics, error)
}

type DiskCollector interface {
	Collect(ctx context.Context) (domain.DiskMetrics, error)
}

type NetworkCollector interface {
	Collect(ctx context.Context) (domain.NetworkMetrics, error)
}
