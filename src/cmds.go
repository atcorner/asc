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
)

// GetEximMailQue
// Returns number of mails in queve
func GetEximMailQue() (que uint64, err error) {

	command := exec.Command("exim", "-bpc")
	defer command.Wait()

	stdout, err := command.StdoutPipe()
	if err != nil {
		return 0, nil
	}

	defer stdout.Close()

	err = command.Start()
	if err != nil {
		return 0, nil
	}

	in := bufio.NewScanner(stdout)
	for in.Scan() {
		que, err = strconv.ParseUint(in.Text(), 10, 64)
		if err != nil {
			return 0, nil
		}
	}

	return que, nil
}
