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
	"os/exec"
	"strconv"
	"strings"
)

// GetTopProcesses ...
// Get top 5 processes by CPU usage
func GetTopProcesses(processes *Processes_t) error {
	proc := Process_t{}

	// Empty slice
	*processes = nil

	command := exec.Command("ps",
		"-Ao",
		"pid,pcpu,rss,user,comm",
		"-r",
	)
	defer command.Wait()

	stdout, err := command.StdoutPipe()
	if err != nil {
		return err
	}

	defer stdout.Close()

	err = command.Start()
	if err != nil {
		return err
	}

	in := bufio.NewScanner(stdout)
	in.Scan()
	for i := uint8(0); i < PROCNUM && in.Scan(); i++ {
		fields := strings.Fields(in.Text())

		proc.Pid, err = strconv.ParseUint(fields[0], 10, 64)
		proc.CPU, err = strconv.ParseFloat(fields[1], 10)
		proc.RSS, err = strconv.ParseUint(fields[2], 10, 64)
		proc.User = fields[3]
		proc.CMD = fields[4]

		// On Linux RSS from ps in in kB
		proc.RSS *= 1024

		*processes = append(*processes, proc)
	}

	return nil
}
