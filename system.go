package main

import (
	"strings"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type SystemMetrics struct {
	CPUUser   float64 `json:"sys.cpu.all.user"`
	CPUNice   float64 `json:"sys.cpu.all.nice"`
	CPUSystem float64 `json:"sys.cpu.all.system"`
	CPUIowait float64 `json:"sys.cpu.all.iowait"`

	Load1m float64 `json:"sys.load.1m"`

	MemUsed        uint64  `json:"sys.memory.mem.used"`
	MemUsedPercent float64 `json:"sys.memory.mem.used.percent"`
	MemTotal       uint64  `json:"sys.memory.mem.total"`

	SwapUsed        uint64  `json:"sys.memory.swap.used"`
	SwapUsedPercent float64 `json:"sys.memory.swap.used.percent"`
	SwapTotal       uint64  `json:"sys.memory.swap.total"`

	NetBytesRecv uint64 `json:"net.bytes.recv"`
	NetBytesSent uint64 `json:"net.bytes.sent"`

	DisksIO    float64 `json:"disks.io"`
	DisksUsage float64 `json:"disks.usage"`
}

/*type FullMetrics struct {
	CPU        cpu.TimesStat
	Load       *load.AvgStat

	Mem        *mem.VirtualMemoryStat
	Swap       *mem.SwapMemoryStat

	Net        []net.IOCountersStat
	Conn       []net.FilterStat

	DisksIO    []disk.IOCountersStat
	Partitions []*disk.UsageStat
}*/

func GetSystemMetrics() SystemMetrics {
	cpus, _ := cpu.Times(false)
	loadm, _ := load.Avg()

	ram, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()

	netio, _ := net.IOCounters(false)
	//netf, _ := net.FilterCounters()

	dss, _ := disk.IOCounters()
	for _, ds := range dss {
		// FIXME: filter 3 letters disks for now ; but what do we cant here?
		if len(ds.Name) == 3 {
			m.DisksIO = append(m.DisksIO, ds)
		}
	}
	pts, _ := disk.Partitions(false)
	for _, pt := range pts {
		if !strings.Contains(pt.Mountpoint, "docker") {
			u, _ := disk.Usage(pt.Mountpoint)
			m.Partitions = append(m.Partitions, u)
		}
	}

	return SystemMetrics{
		CPUUser:         cpus[0].User,
		CPUSystem:       cpus[0].System,
		CPUIowait:       cpus[0].Iowait,
		Load1m:          loadm.Load1,
		MemUsed:         ram.Used,
		MemUsedPercent:  ram.UsedPercent,
		MemTotal:        ram.Total,
		SwapUsed:        swap.Used,
		SwapUsedPercent: swap.UsedPercent,
		SwapTotal:       swap.Total,
		NetBytesRecv:    netio[0].BytesRecv,
		NetBytesSent:    netio[0].BytesSent,
		//DisksUsage:     0,
		//DisksIO:         0,
	}
}
