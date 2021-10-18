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
#include <mach/mach_host.h>
*/
import "C"

import (
	"bytes"
	"encoding/binary"
	"errors"
	"syscall"
	"unsafe"
)

// generic Sysctl buffer unmarshalling
func sysctlbyname(name string, data interface{}) (err error) {
	val, err := syscall.Sysctl(name)

	if err != nil {
		return err
	}

	buf := []byte(val)

	switch v := data.(type) {
	case *uint64:
		*v = *(*uint64)(unsafe.Pointer(&buf[0]))
		return
	}

	bbuf := bytes.NewBuffer([]byte(val))

	return binary.Read(bbuf, binary.LittleEndian, data)
}

// get memory info
func vmInfo(vmstat *C.vm_statistics_data_t) error {
	var count C.mach_msg_type_number_t = C.HOST_VM_INFO_COUNT

	status := C.host_statistics(
		C.host_t(C.mach_host_self()),
		C.HOST_VM_INFO,
		C.host_info_t(unsafe.Pointer(vmstat)),
		&count)

	if status != C.KERN_SUCCESS {
		return errors.New("host_statistics error")
	}

	return nil
}
