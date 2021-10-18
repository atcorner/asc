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
	"io/ioutil"
)

const HTTP_PORT_DEFAULT = 8080
const COLLECT_INTERVAL_DEFAULT = 10
const COLLECT_INTERVAL_MIN = 10
const BUFFER_SIZE_DEFAULT = 4
const BUFFER_SIZE_MIN = 2

type Config_t struct {
	HTTPPort        uint16  `json:"http_port"`
	CollectInterval uint8   `json:"collect_interval"`
	BufferSize      uint8   `json:"buffer_size"`
	Disks           Disks_t `json:"disks"`
}

type DiskConfig_t struct {
	MountPoint string `json:"mount_point"`
}

func GetConfiguration(fileName string, config *Config_t) error {

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, config)
	if err != nil {
		return err
	}

	// Set minimum values
	if config.HTTPPort < 1024 {
		config.HTTPPort = HTTP_PORT_DEFAULT
	}

	if config.CollectInterval < 5 {
		config.CollectInterval = COLLECT_INTERVAL_MIN
	}

	if config.BufferSize < 2 {
		config.BufferSize = BUFFER_SIZE_MIN
	}

	return nil
}

func GetConfigurationDefaults(config *Config_t) {
	config.HTTPPort = HTTP_PORT_DEFAULT
	config.CollectInterval = COLLECT_INTERVAL_DEFAULT
	config.BufferSize = BUFFER_SIZE_DEFAULT
}
