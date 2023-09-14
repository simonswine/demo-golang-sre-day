package main

import (
	"fmt"
	"net/http"
	"time"
)

func sleepHandler(w http.ResponseWriter, r *http.Request) {
	sleep := r.URL.Query().Get("sleep")
	sleepD, err := time.ParseDuration(sleep)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "bad duration: %s: %s\n", sleep, err)
	}
	cpuIntensiveTask()
	time.Sleep(sleepD)
	fmt.Fprintf(w, "slept for: %s\n", sleepD)
}
