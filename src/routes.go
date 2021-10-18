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
	"github.com/gorilla/mux"
)

func NewRouter(context *appCtx_t) *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.
		Methods("GET").
		Name("Index").
		Path("/").
		Handler(handler_t{context, HIndex})

	router.
		Methods("GET").
		Name("CollectAll").
		Path("/Collect/All").
		Handler(handler_t{context, HCollectAll})

	router.
		Methods("GET").
		Name("Memstats").
		Path("/Collect/Memstats").
		Handler(handler_t{context, HMemStats})

	return router
}
