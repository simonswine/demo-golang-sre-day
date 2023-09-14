package main

import (
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	_ "net/http/pprof"

	"github.com/felixge/fgprof"
)

const (
	sleepTime   = 10 * time.Millisecond
	cpuTime     = 30 * time.Millisecond
	networkTime = 60 * time.Millisecond
)

func main() {
	// Run http endpoints for both pprof and fgprof.
	http.DefaultServeMux.HandleFunc("/sleep", sleepHandler)
	http.DefaultServeMux.Handle("/debug/fgprof", fgprof.Handler())
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		addr := "0.0.0.0:6060"
		log.Printf("Listening on %s", addr)
		log.Println(http.ListenAndServe(addr, nil))
	}()

	// clients run request in a loop
	if url := os.Getenv("SERVER_URL"); url != "" {
		for {
			// Http request to a web service that might be slow.
			slowNetworkRequest(url)
			// Some heavy CPU computation.
			cpuIntensiveTask()
			// Poorly named function that you don't understand yet.
			weirdFunction()
		}
	}

	// wait for the http to terminate
	wg.Wait()
}

func slowNetworkRequest(url string) {
	res, err := http.Get(url + "?sleep=" + networkTime.String())
	if err != nil {
		log.Println("error ", err)
		return
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Printf("error bad code: %d\n", res.StatusCode)
		return
	}
}

func cpuIntensiveTask() {
	start := time.Now()
	for time.Since(start) <= cpuTime {
		// Spend some time in a hot loop to be a little more realistic than
		// spending all time in time.Since().
		for i := 0; i < 1000; i++ {
			_ = i
		}
	}
}

func weirdFunction() {
	time.Sleep(sleepTime)
}
