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
#include <mach/vm_map.h>
*/
import "C"

// GetMemInfo ...
// Get memory info
func GetMemInfo(mem *Memory_t) error {
	s := MemSample_t{}

	err := getMemSample(&s)
	if err != nil {
		return err
	}

	mem.RamSize = s.MemTotal
	mem.SwapSize = s.SwapTotal

	// To avoid division by zero we will add 1 byte to total size of ram & swap
	s.MemTotal++
	s.SwapTotal++

	ramUsed := s.MemTotal - s.MemFree - s.Cached - s.Buffers

	mem.RamUtilization = uint8(100 * ramUsed / s.MemTotal)
	mem.CacheUtilization = uint8(100 * s.Cached / s.MemTotal)
	mem.BuffersUtilization = uint8(100 * s.Buffers / s.MemTotal)
	mem.SwapUtilization = uint8(100 * (s.SwapTotal - s.SwapFree) / s.SwapTotal)

	return nil
}

type xswUsage struct {
	Total, Avail, Used uint64
}

func getMemSample(sample *MemSample_t) error {
	var vmstat C.vm_statistics_data_t

	var totalMemory, kern, free, actualFree uint64

	if err := sysctlbyname("hw.memsize", &totalMemory); err != nil {
		return err
	}

	if err := vmInfo(&vmstat); err != nil {
		return err
	}

	kern = uint64(vmstat.inactive_count) << 12
	free = uint64(vmstat.free_count) << 12

	actualFree = free + kern

	swUsage := xswUsage{}

	if err := sysctlbyname("vm.swapusage", &swUsage); err != nil {
		return err
	}

	sample.MemTotal = totalMemory
	sample.MemFree = actualFree
	sample.Buffers = kern
	sample.Cached = 0
	sample.SwapTotal = swUsage.Total
	sample.SwapFree = swUsage.Avail

	return nil
}
