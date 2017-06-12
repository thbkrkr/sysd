package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
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

	NetConnTrackCount int64 `json:"net.conntrack.count"`
	NetConnTrackMax   int64 `json:"net.conntrack.max"`

	//DisksIO    float64 `json:"disks.io"`
	//DisksUsage float64 `json:"disks.usage"`
}

func GetSystemMetrics() SystemMetrics {
	//cpus, _ := cpu.Times(false)
	loadm, _ := load.Avg()

	ram, _ := mem.VirtualMemory()
	swap, _ := mem.SwapMemory()

	netf, _ := net.FilterCounters()

	/*dss, _ := disk.IOCounters()
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
	}*/

	return SystemMetrics{
		/*	CPUUser:         cpus[0].User,
			CPUSystem:       cpus[0].System,
			CPUIowait:       cpus[0].Iowait,*/
		Load1m:          loadm.Load1,
		MemUsed:         ram.Used,
		MemUsedPercent:  ram.UsedPercent,
		MemTotal:        ram.Total,
		SwapUsed:        swap.Used,
		SwapUsedPercent: swap.UsedPercent,
		SwapTotal:       swap.Total,
		/*		NetBytesRecv:      netio[0].BytesRecv,
				NetBytesSent:      netio[0].BytesSent,*/
		NetConnTrackCount: netf[0].ConnTrackCount,
		NetConnTrackMax:   netf[0].ConnTrackMax,
		//DisksUsage:     0,
		//DisksIO:         0,
	}
}

var now = fmt.Sprintf("%v ", time.Now().UnixNano()/1000)

func GetCPUMetrics() string {
	ts := ""

	cpuTimes, _ := cpu.Times(false)
	timeStat := cpuTimes[0]

	totalCPUTime := timeStat.Total()
	totalCPUUsage := totalCPUTime - (timeStat.Idle + timeStat.Iowait)
	totalCPUPct := totalCPUUsage / totalCPUTime * 100

	ts += fmt.Sprintf("%v os.cpu.usage{} %v\n", now, totalCPUPct)

	/*`"user":` + strconv.FormatFloat(c.User, 'f', 1, 64),
	`"system":` + strconv.FormatFloat(c.System, 'f', 1, 64),
	`"idle":` + strconv.FormatFloat(c.Idle, 'f', 1, 64),
	`"nice":` + strconv.FormatFloat(c.Nice, 'f', 1, 64),
	`"iowait":` + strconv.FormatFloat(c.Iowait, 'f', 1, 64),
	`"irq":` + strconv.FormatFloat(c.Irq, 'f', 1, 64),
	`"softirq":` + strconv.FormatFloat(c.Softirq, 'f', 1, 64),
	`"steal":` + strconv.FormatFloat(c.Steal, 'f', 1, 64),
	`"guest":` + strconv.FormatFloat(c.Guest, 'f', 1, 64),
	`"guestNice":` + strconv.FormatFloat(c.GuestNice, 'f', 1, 64),
	`"stolen":` + strconv.FormatFloat(c.Stolen, 'f', 1, 64),
	*/
	return ts
}

func GetNetMetrics() string {
	counters, _ := net.IOCounters(false)
	ts := ""
	for _, cnt := range counters {
		ts += fmt.Sprintf("%v os.net.bytes{iface=%v,direction=in} %v\n", now, cnt.Name, cnt.BytesRecv)
		ts += fmt.Sprintf("%v os.net.bytes{iface=%v,direction=out} %v\n", now, cnt.Name, cnt.BytesSent)
		ts += fmt.Sprintf("%v os.net.packets{iface=%v,direction=in} %v\n", now, cnt.Name, cnt.PacketsRecv)
		ts += fmt.Sprintf("%v os.net.packets{iface=%v,direction=out} %v\n", now, cnt.Name, cnt.PacketsSent)
		ts += fmt.Sprintf("%v os.net.errs{iface=%v,direction=in} %v\n", now, cnt.Name, cnt.Errin)
		ts += fmt.Sprintf("%v os.net.errs{iface=%v,direction=out} %v\n", now, cnt.Name, cnt.Errout)
		ts += fmt.Sprintf("%v os.net.dropped{iface=%v,direction=in} %v\n", now, cnt.Name, cnt.Dropin)
		ts += fmt.Sprintf("%v os.net.dropped{iface=%v,direction=out} %v\n", now, cnt.Name, cnt.Dropout)
	}
	return ts
}
