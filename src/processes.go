package main

const PROCNUM uint8 = 5

type Process_t struct {
	Pid  uint64  `json:"pid"`
	CPU  float64 `json:"cpu"`
	RSS  uint64  `json:"rss"`
	User string  `json:"user"`
	CMD  string  `json:"cmd"`
}

type Processes_t []Process_t
