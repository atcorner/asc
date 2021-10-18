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
	"golang.org/x/sys/unix"
)

// Sets system uptime and load averages
func GetUptime(system *SystemInfo_t) error {
	s := unix.Sysinfo_t{}
	err := unix.Sysinfo(&s)
	if err != nil {
		return err
	}

	system.Uptime = s.Uptime

	for i := range s.Loads {
		// s.Loads are left bit shifted by SI_LOAD_SHIFT
		// /usr/src/linux/include/linux/kernel.h: #define SI_LOAD_SHIFT 16
		system.Loads[i] = float32((s.Loads[i]*100)>>16) / 100
	}

	return nil
}
