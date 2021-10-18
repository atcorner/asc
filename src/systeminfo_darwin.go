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
*/
import "C"

import (
	"syscall"
	"time"
)

// GetUptime ...
// Sets system uptime and load averages
func GetUptime(system *SystemInfo_t) error {
	tv := syscall.Timeval32{}

	if err := sysctlbyname("kern.boottime", &tv); err != nil {
		return err
	}

	system.Uptime = int64(time.Since(time.Unix(int64(tv.Sec), int64(tv.Usec)*1000)).Seconds())

	avg := []C.double{0, 0, 0}

	C.getloadavg(&avg[0], C.int(len(avg)))

	system.Loads[0] = float32(avg[0])
	system.Loads[1] = float32(avg[1])
	system.Loads[2] = float32(avg[2])

	return nil
}
