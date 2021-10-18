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
	"log"
	"net/http"
	"runtime"
	"strconv"
	"time"
)

type appCtx_t struct {
	Config    Config_t
	Buffer    Buffer_t
	Channel   chan Buffer_t
	MemStats  runtime.MemStats
	StartTime int64
}

func main() {
	log.Println("Starting up ...")
	ctx := &appCtx_t{}

	log.Println("Loading configuration ...")
	err := GetConfiguration("config.json", &ctx.Config)
	if err != nil {
		GetConfigurationDefaults(&ctx.Config)
		log.Printf("Error: %v. Using default configuration.\n", err)
		log.Printf("http_port = %v\n", ctx.Config.HTTPPort)
		log.Printf("collect_interval = %v\n", ctx.Config.CollectInterval)
		log.Printf("message_buffer_size = %v\n", ctx.Config.BufferSize)
	}

	ctx.Channel = make(chan Buffer_t, ctx.Config.BufferSize)
	ctx.StartTime = time.Now().Unix()

	log.Println("Starting collector thread ...")
	go StartCollector(ctx)

	router := NewRouter(ctx)
	port := strconv.Itoa(int(ctx.Config.HTTPPort))
	log.Println("Starting REST server ...")
	log.Fatal(http.ListenAndServe(":"+port, router))
}
