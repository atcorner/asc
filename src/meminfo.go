package main

type Memory_t struct {
	RamSize            uint64 `json:"ram_size"`
	RamUtilization     uint8  `json:"ram_utilization"`
	CacheUtilization   uint8  `json:"cache_utilization"`
	BuffersUtilization uint8  `json:"buffers_utilization"`
	SwapSize           uint64 `json:"swap_size"`
	SwapUtilization    uint8  `json:"swap_utilization"`
}

type MemSample_t struct {
	MemTotal  uint64 `json:"mem_total"`
	MemFree   uint64 `json:"mem_free"`
	Buffers   uint64 `json:"buffers"`
	Cached    uint64 `json:"cached"`
	SwapTotal uint64 `json:"swap_total"`
	SwapFree  uint64 `json:"swap_free"`
}
