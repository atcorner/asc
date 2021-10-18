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

func getDiskUtilization(disk *Disk_t) error {
	buf := unix.Statfs_t{}
	err := unix.Statfs(disk.MountPoint, &buf)
	if err != nil {
		return err
	}

	size := uint64(buf.Bsize) * uint64(buf.Blocks)
	free := uint64(buf.Bsize) * uint64(buf.Bavail)
	utilization := uint8(100 * (size - free) / size)

	disk.Size = size
	disk.Utilization = utilization

	return nil
}
