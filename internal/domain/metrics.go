package domain

import "time"

type SystemMetrics struct {
	Timestamp time.Time
	CPU       CPUMetrics
	Memory    MemoryMetrics
	Disk      DiskMetrics
	Network   NetworkMetrics
}

type CPUMetrics struct {
	TotalUsage  float64   // Общая загрузка в %
	CoreUsages  []float64 // Загрузка по ядрам в %
	Temperature float64   // Температура в °C
}

type MemoryMetrics struct {
	Total       uint64  // Общий объем в байтах
	Used        uint64  // Используемый объем в байтах
	UsedPercent float64 // Процент использования
	SwapTotal   uint64  // Общий объем swap в байтах
	SwapUsed    uint64  // Используемый объем swap в байтах
}

type DiskMetrics struct {
	Partitions []DiskPartition   // Информация о разделах
	IOCounters map[string]DiskIO // Статистика ввода-вывода по устройствам
}

type DiskPartition struct {
	Device      string  // Устройство
	Mountpoint  string  // Точка монтирования
	Total       uint64  // Общий объем в байтах
	Used        uint64  // Используемый объем в байтах
	UsedPercent float64 // Процент использования
}

type DiskIO struct {
	ReadCount  uint64 // Количество операций чтения
	WriteCount uint64 // Количество операций записи
	ReadBytes  uint64 // Прочитано байт
	WriteBytes uint64 // Записано байт
}

type NetworkMetrics struct {
	Interfaces map[string]NetworkInterface // Статистика по интерфейсам
}

type NetworkInterface struct {
	BytesSent   uint64 // Отправлено байт
	BytesRecv   uint64 // Получено байт
	PacketsSent uint64 // Отправлено пакетов
	PacketsRecv uint64 // Получено пакетов
}
