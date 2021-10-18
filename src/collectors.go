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
	"encoding/json"
	"log"
	"time"
)

type Collector_t struct {
	Timestamp  time.Time    `json:"timestamp"`
	SystemInfo SystemInfo_t `json:"system"`
}

type Buffer_t struct {
	Buf []byte
	Err error
}

func StartCollector(app *appCtx_t) {
	var collector Collector_t
	var buffer Buffer_t

	// Get disks to collect info from configuration
	for i := range app.Config.Disks {
		collector.SystemInfo.Disks = append(
			collector.SystemInfo.Disks,
			app.Config.Disks[i],
		)
	}

	// Get *all* system info outside of loop
	// inside loop we will only update changeable data
	buffer.Err = GetSystemInfo(&collector.SystemInfo)
	if buffer.Err != nil {
		log.Println(buffer.Err)
	}

	for {
		collector.Timestamp = time.Now().UTC()

		buffer.Buf, _ = json.Marshal(collector)
		select {
		case app.Channel <- buffer:
		default:
			// Channel is full, drain it
			<-app.Channel
		}

		// Substract spent sleep time from GetCPUUtilization
		time.Sleep(time.Duration(app.Config.CollectInterval-3) * time.Second)

		buffer.Err = UpdateSystemInfo(&collector.SystemInfo)
		if buffer.Err != nil {
			log.Println(buffer.Err)
		}
	}
}
