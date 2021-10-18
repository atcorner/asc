/*
 * This Source Code Form is subject to the terms of the Mozilla Public License,
 * v. 2.0. If a copy of the MPL was not distributed with this file, You can
 * obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Copyright 2016, Ante Vojvodic, <ante@atcorner.hr>
 * All rights reserved.
 */

package main

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"strconv"
	"strings"
)

const PROCMEMINFO string = "/proc/meminfo"

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

	ram_used := s.MemTotal - s.MemFree - s.Cached - s.Buffers

	mem.RamUtilization = uint8(100 * ram_used / s.MemTotal)
	mem.CacheUtilization = uint8(100 * s.Cached / s.MemTotal)
	mem.BuffersUtilization = uint8(100 * s.Buffers / s.MemTotal)
	mem.SwapUtilization = uint8(100 * (s.SwapTotal - s.SwapFree) / s.SwapTotal)

	return nil
}

func getMemSample(sample *MemSample_t) error {
	// While the file shows kilobytes (kB; 1 kB equals 1000 B), it is actually
	// kibibytes (KiB; 1 KiB equals 1024 B). This imprecision in /proc/meminfo
	// is known, but is not corrected due to legacy concerns - programs rely on
	// /proc/meminfo to specify size with the "kB" string.
	m := uint64(1024) //Memory unit size in bytes

	file, err := ioutil.ReadFile(PROCMEMINFO)
	if err != nil {
		return err
	}

	in := bufio.NewScanner(bytes.NewBuffer(file))
	for in.Scan() {
		fields := strings.Fields(in.Text())

		if len(fields) == 3 {
			val, err := strconv.ParseUint(fields[1], 10, 64)
			if err != nil {
				return err
			}
			switch fields[0] {
			case "MemTotal:":
				sample.MemTotal = val * m
			case "MemFree:":
				sample.MemFree = val * m
			case "Buffers:":
				sample.Buffers = val * m
			case "Cached:":
				sample.Cached = val * m
			case "SwapTotal:":
				sample.SwapTotal = val * m
			case "SwapFree:":
				sample.SwapFree = val * m
			}
		}
	}

	return nil
}
