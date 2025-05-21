package metrics

import (
	"context"
	"github.com/shirou/gopsutil/v3/net"
	"monitoring/internal/domain"
)

type GopsutilNetworkCollector struct{}

func NewGopsutilNetworkCollector() *GopsutilNetworkCollector {
	return &GopsutilNetworkCollector{}
}

func (nc *GopsutilNetworkCollector) Collect(ctx context.Context) (domain.NetworkMetrics, error) {
	ioCounters, err := net.IOCountersWithContext(ctx, true)
	if err != nil {
		return domain.NetworkMetrics{}, err
	}

	interfaces := make(map[string]domain.NetworkInterface)
	for _, iface := range ioCounters {
		interfaces[iface.Name] = domain.NetworkInterface{
			BytesSent:   iface.BytesSent,
			BytesRecv:   iface.BytesRecv,
			PacketsSent: iface.PacketsSent,
			PacketsRecv: iface.PacketsRecv,
		}
	}

	return domain.NetworkMetrics{
		Interfaces: interfaces,
	}, nil
}
