package metrics

import (
	"context"
	"github.com/shirou/gopsutil/v3/cpu"
	"monitoring/internal/domain"
)

type GopsutilCPUCollector struct{}

func NewGopsutilCollector() *GopsutilCPUCollector {
	return &GopsutilCPUCollector{}
}

func (cc *GopsutilCPUCollector) Collect(ctx context.Context) (domain.CPUMetrics, error) {
	percentages, err := cpu.PercentWithContext(ctx, 0, true)
	if err != nil {
		return domain.CPUMetrics{}, err
	}
	totalPercent, err := cpu.PercentWithContext(ctx, 0, false)
	if err != nil {
		return domain.CPUMetrics{}, err
	}
	return domain.CPUMetrics{
		TotalUsage: totalPercent[0],
		CoreUsages: percentages,
	}, nil
}
