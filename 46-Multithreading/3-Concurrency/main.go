package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number uint64 = 0

func main() {
	// m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// m.Lock()
		atomic.AddUint64(&number, 1)
		// number++
		// m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("You are visitant number %d", number)))
	})
	http.ListenAndServe(":3000", nil)
}
