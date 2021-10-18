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
	"os"
	"runtime"
)

type SystemInfo_t struct {
	Hostname     string      `json:"hostname"`
	IPAddress    string      `json:"ip_address"`
	Uptime       int64       `json:"uptime"`
	Loads        [3]float32  `json:"loads"`
	CPU          CPUInfo_t   `json:"cpu"`
	Memory       Memory_t    `json:"memory"`
	Disks        Disks_t     `json:"disks"`
	TopProcesses Processes_t `json:"top_processes"`
	EximMailQue  uint64      `json:"exim_mail_que"`
}

func GetSystemInfo(s *SystemInfo_t) error {
	var err error = nil

	s.Hostname, err = os.Hostname()
	if err != nil {
		return err
	}

	err = GetHostnameIPAddress(s)
	if err != nil {
		return err
	}

	err = GetUptime(s)
	if err != nil {
		return err
	}

	s.CPU.Count = uint16(runtime.NumCPU())

	err = GetCPUUtilization(&s.CPU.Utilization)
	if err != nil {
		return err
	}

	err = GetMemInfo(&s.Memory)
	if err != nil {
		return err
	}

	err = GetDisksInfo(&s.Disks)
	if err != nil {
		return err
	}

	err = GetTopProcesses(&s.TopProcesses)
	if err != nil {
		return err
	}

	s.EximMailQue, err = GetEximMailQue()
	if err != nil {
		return err
	}

	return nil
}

func UpdateSystemInfo(s *SystemInfo_t) error {
	var err error = nil

	err = GetUptime(s)
	if err != nil {
		return err
	}

	err = GetCPUUtilization(&s.CPU.Utilization)
	if err != nil {
		return err
	}

	err = GetMemInfo(&s.Memory)
	if err != nil {
		return err
	}

	err = GetDisksInfo(&s.Disks)
	if err != nil {
		return err
	}

	err = GetTopProcesses(&s.TopProcesses)
	if err != nil {
		return err
	}

	s.EximMailQue, err = GetEximMailQue()
	if err != nil {
		return nil
	}
	return nil
}
