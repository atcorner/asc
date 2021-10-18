package main

type CPUInfo_t struct {
	Count       uint16           `json:"count"`
	Utilization CPUUtilization_t `json:"utilization"`
}

type CPUUtilization_t struct {
	User   uint8 `json:"user"`
	System uint8 `json:"system"`
	IOWait uint8 `json:"io_wait"`
}
