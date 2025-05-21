package metrics

import (
	"context"
	"github.com/shirou/gopsutil/v3/mem"
	"monitoring/internal/domain"
)

type GopsutilMemoryCollector struct {
}

func NewGopsutillMemoryCollector() *GopsutilMemoryCollector {
	return &GopsutilMemoryCollector{}
}

func (mc *GopsutilMemoryCollector) Collect(ctx context.Context) (domain.MemoryMetrics, error) {
	v, err := mem.VirtualMemoryWithContext(ctx)
	if err != nil {
		return domain.MemoryMetrics{}, err
	}

	s, err := mem.SwapMemoryWithContext(ctx)
	if err != nil {
		return domain.MemoryMetrics{}, err
	}

	return domain.MemoryMetrics{
		Total:       v.Total,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
		SwapTotal:   s.Total,
		SwapUsed:    s.Used,
	}, nil
}
