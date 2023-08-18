package utils

import (
	"encoding/json"
	"gokube/stats"
	"net/http"
	"net/http/httptest"

	"github.com/c9s/goprocinfo/linux"
)

func GetStats(diskTotal int, diskFree int, memTotal int, memUsed int, user uint64, nice uint64, sys uint64, idle uint64, iowait uint64, irq uint64, softirq uint64, steal uint64, guest uint64, guest_nice uint64) *stats.Stats {
	memstats := linux.MemInfo{
		MemTotal:     uint64(memTotal),
		MemAvailable: uint64(memTotal - memUsed),
	}

	diskstats := linux.Disk{
		All:  uint64(diskTotal),
		Free: uint64(diskFree),
	}

	cpustats := linux.CPUStat{
		User:      user,
		Nice:      nice,
		System:    sys,
		Idle:      idle,
		IOWait:    iowait,
		IRQ:       irq,
		SoftIRQ:   softirq,
		Steal:     steal,
		Guest:     guest,
		GuestNice: guest_nice,
	}

	//cpuinfo := linux.CPUInfo{
	//	Processors: []linux.Processor{{
	//		Id:       0,
	//		VendorId: "GenuineIntel",
	//		Cores:    6,
	//	}},
	//}

	s := stats.Stats{
		MemStats:  &memstats,
		DiskStats: &diskstats,
		CpuStats:  &cpustats,
		//CpuInfo:   &cpuinfo,
	}

	return &s
}

func CreateTestServer(returnData interface{}) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(returnData)
	}))
	return ts
}
