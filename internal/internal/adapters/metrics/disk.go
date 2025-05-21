package metrics

import (
	"context"
	"github.com/shirou/gopsutil/v3/disk"
	"monitoring/internal/domain"
)

type GopsutilDiskCollector struct{}

func NewGopsutilDiskCollector() *GopsutilDiskCollector {
	return &GopsutilDiskCollector{}
}
func (dc *GopsutilDiskCollector) Collect(ctx context.Context) (domain.DiskMetrics, error) {
	partitions, err := disk.PartitionsWithContext(ctx, false)
	if err != nil {
		return domain.DiskMetrics{}, err
	}
	var diskPartitions []domain.DiskPartition
	for _, p := range partitions {
		usage, err := disk.UsageWithContext(ctx, p.Mountpoint)
		if err != nil {
			continue
		}

		diskPartitions = append(diskPartitions, domain.DiskPartition{
			Device:      p.Device,
			Mountpoint:  p.Mountpoint,
			Total:       usage.Total,
			Used:        usage.Used,
			UsedPercent: usage.UsedPercent,
		})
	}
	ioCounters, err := disk.IOCountersWithContext(ctx)
	if err != nil {
		return domain.DiskMetrics{
			Partitions: diskPartitions,
			IOCounters: make(map[string]domain.DiskIO),
		}, nil
	}

	diskIOMap := make(map[string]domain.DiskIO)
	for name, counter := range ioCounters {
		diskIOMap[name] = domain.DiskIO{
			ReadCount:  counter.ReadCount,
			WriteCount: counter.WriteCount,
			ReadBytes:  counter.ReadBytes,
			WriteBytes: counter.WriteBytes,
		}
	}

	return domain.DiskMetrics{
		Partitions: diskPartitions,
		IOCounters: diskIOMap,
	}, nil
}
