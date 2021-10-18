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
	"fmt"
	"net/http"
	"runtime"
	"time"
)

type handler_t struct {
	*appCtx_t
	H func(*appCtx_t, http.ResponseWriter, *http.Request)
}

func (ah handler_t) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ah.H(ah.appCtx_t, w, r)
}

func HIndex(a *appCtx_t, w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ATC STATUS COLLECTOR")
}

func HCollectAll(a *appCtx_t, w http.ResponseWriter, r *http.Request) {
	select {
	case a.Buffer = <-a.Channel:
	default:
		// Channel is empty
		if len(a.Buffer.Buf) == 0 {
			// On first run Buffer is empty, wait for collector
			a.Buffer = <-a.Channel
		}
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")

	if a.Buffer.Err != nil {
		http.Error(w, a.Buffer.Err.Error(), 500)
	} else {
		fmt.Fprintf(w, "%s", string(a.Buffer.Buf))
	}
}

func HMemStats(a *appCtx_t, w http.ResponseWriter, r *http.Request) {
	runtime.ReadMemStats(&a.MemStats)

	fmt.Fprintf(w,
		"%v %v %v %v %v %v %v %v %v\n",
		time.Now().Unix()-a.StartTime,
		a.MemStats.Alloc/1024,
		a.MemStats.TotalAlloc/1024,
		a.MemStats.HeapAlloc/1024,
		a.MemStats.HeapSys/1024,
		a.MemStats.HeapReleased/1024,
		a.MemStats.HeapIdle/1024,
		a.MemStats.Mallocs,
		a.MemStats.Frees,
	)
}
