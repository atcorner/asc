/*
 * This Source Code Form is subject to the terms of the Mozilla Public License,
 * v. 2.0. If a copy of the MPL was not distributed with this file, You can
 * obtain one at http://mozilla.org/MPL/2.0/.
 *
 * Copyright 2016, Ante Vojvodic, <ante@atcorner.hr>
 * All rights reserved.
 */

package main

type Disk_t struct {
	MountPoint  string `json:"mount_point"` // Disk mount point
	Size        uint64 `json:"size"`        // Size in bytes
	Utilization uint8  `json:"utilization"` // Disk utilization in percents
}

type Disks_t []Disk_t

func GetDisksInfo(disks *Disks_t) error {
	for i := range *disks {
		err := GetDiskInfo(&(*disks)[i])
		if err != nil {
			return err
		}
	}

	return nil
}

func GetDiskInfo(disk *Disk_t) error {
	err := getDiskUtilization(disk)
	if err != nil {
		return err
	}

	return nil
}
