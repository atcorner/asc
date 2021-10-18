/*
 * This Source Code Form is subject to the terms of the Mozilla Public License,
 * v. 2.0. If a copy of the MPL was not distributed with this file, You can
 * obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Copyright 2016, Ante Vojvodic, <ante@atcorner.hr>
 * All rights reserved.
 */

package main

/*
#include <stdlib.h>
#include <mach/mach_host.h>
*/
import "C"

import (
	"errors"
	"time"
	"unsafe"
)

// GetCPUUtilization ...
// Get average CPU utilization for 3 seconds
func GetCPUUtilization(cpu *CPUUtilization_t) error {
	u1, s1, i1, err := _GetTicks()
	if err != nil {
		return err
	}
	time.Sleep(3 * time.Second)
	u2, s2, i2, err := _GetTicks()
	if err != nil {
		return err
	}

	du := u2 - u1
	ds := s2 - s1
	di := i2 - i1

	total := du + ds + di

	var onePercent float32
	onePercent = float32(total) / 100.0

	cpu.User = uint8(float32(du) / onePercent)
	cpu.System = uint8(float32(ds) / onePercent)
	cpu.IOWait = uint8(float32(di) / onePercent)
	return nil
}

func _GetTicks() (user, system, iowait uint64, error error) {
	var count C.mach_msg_type_number_t = C.HOST_CPU_LOAD_INFO_COUNT
	var cpuload C.host_cpu_load_info_data_t

	status := C.host_statistics(C.host_t(C.mach_host_self()),
		C.HOST_CPU_LOAD_INFO,
		C.host_info_t(unsafe.Pointer(&cpuload)),
		&count)

	if status != C.KERN_SUCCESS {
		return 0, 0, 0, errors.New("Error getting ticks")
	}

	user = uint64(cpuload.cpu_ticks[C.CPU_STATE_USER])
	system = uint64(cpuload.cpu_ticks[C.CPU_STATE_SYSTEM] + cpuload.cpu_ticks[C.CPU_STATE_NICE])
	iowait = 0
	return
}
