package main

import (
	"net/http"
	"time"
)

func writeEmoji(w http.ResponseWriter, req *http.Request) {
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	f, ok := w.(http.Flusher)

	if !ok {
		return
	}

	for {
		select {
		case <-req.Context().Done():
			return
		case <-t.C:
			if _, err := w.Write([]byte("💜")); err != nil {
				return
			}
			f.Flush()
		}
	}
}

func main() {
	http.HandleFunc("/", writeEmoji)
	http.ListenAndServe(":8080", nil)
}
