// Copyright 2018 Google LLC. All rights reserved.
// Use of this source code is governed by the Apache 2.0
// license that can be found in the LICENSE file.

// [START functions_helloworld_get]

// Package helloworld provides a set of Cloud Function samples.
package helloworld

import (
	"fmt"
	"net/http"
	"time"
)

func task1(ch chan string) {
	time.Sleep(1 * time.Second)
	ch <- "task1"
}

func task2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "task2"
}

// ProcessQuery is an HTTP Cloud Function.
func ProcessQuery(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string)
	go task1(ch)
	go task2(ch)
	task := <-ch
	fmt.Fprint(w, task)
}

// [END functions_helloworld_get]
