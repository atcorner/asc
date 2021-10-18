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
	"net"
)

func GetHostnameIPAddress(s *SystemInfo_t) error {

	addr, err := net.ResolveIPAddr("", s.Hostname)
	if err != nil {
		return err
	}

	s.IPAddress = addr.IP.String()

	return nil
}
