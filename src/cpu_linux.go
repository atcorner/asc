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
	"time"
)

const PROCSTAT string = "/proc/stat"

type CPUSample_t struct {
	User    uint64 `json:"user"`     // time spent in user mode
	Nice    uint64 `json:"nice"`     // time spent in user mode with low priority (nice)
	System  uint64 `json:"system"`   // time spent in system mode
	Idle    uint64 `json:"idle"`     // time spent in the idle task
	IOWait  uint64 `json:"iowait"`   // time spent waiting for I/O to complete (since Linux 2.5.41)
	Irq     uint64 `json:"irq"`      // time spent servicing  interrupts  (since  2.6.0-test4)
	SoftIrq uint64 `json:"soft_irq"` // time spent servicing softirqs (since 2.6.0-test4)
	Steal   uint64 `json:"steal"`    // time spent in other OSes when running in a virtualized environment
	Guest   uint64 `json:"guest"`    // time spent running a virtual CPU for guest operating systems under the control of the Linux kernel.
	Total   uint64 `json:"total"`    // total of all time fields
}

func GetCPUUtilization(cpu *CPUUtilization_t) error {
	s0 := CPUSample_t{}
	s1 := CPUSample_t{}

	err := getCPUSample(&s0)
	if err != nil {
		return err
	}

	time.Sleep(3 * time.Second)

	err = getCPUSample(&s1)
	if err != nil {
		return err
	}

	q := float64(100 / float64(s1.Total-s0.Total))

	cpu.User = uint8(float64(s1.User-s0.User) * q)
	cpu.System = uint8(float64(s1.System-s0.System) * q)
	cpu.IOWait = uint8(float64(s1.IOWait-s0.IOWait) * q)

	return nil
}

func getCPUSample(sample *CPUSample_t) error {
	file, err := ioutil.ReadFile(PROCSTAT)
	if err != nil {
		return err
	}

	in := bufio.NewScanner(bytes.NewBuffer(file))
	for in.Scan() {
		fields := strings.Fields(in.Text())

		if len(fields) > 0 {
			if fields[0] == "cpu" {
				err := parseCPUFields(fields, sample)
				if err != nil {
					return err
				}
				// We have cpu sample, stop parsig lines
				break
			}
		}
	}

	return nil
}

func parseCPUFields(fields []string, sample *CPUSample_t) error {
	n := len(fields)
	for i := 1; i < n; i++ {
		val, err := strconv.ParseUint(fields[i], 10, 64)
		if err != nil {
			return err
		}

		sample.Total += val
		switch i {
		case 1:
			sample.User = val
		case 2:
			sample.Nice = val
		case 3:
			sample.System = val
		case 4:
			sample.Idle = val
		case 5:
			sample.IOWait = val
		case 6:
			sample.Irq = val
		case 7:
			sample.SoftIrq = val
		case 8:
			sample.Steal = val
		case 9:
			sample.Guest = val
		}
	}

	return nil
}
