package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httputil"
	"os"
)

var listenAddr = flag.String("addr", ":4000", "HTTP network address")

func main() {
	flag.Parse()
	mux := http.NewServeMux()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	mux.HandleFunc("/", home)

	logger.Info("starting server", "addr", *listenAddr)

	err := http.ListenAndServe(*listenAddr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}

func home(w http.ResponseWriter, r *http.Request) {
	flag.Parse()
	repDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		return
	}

	host, _ := os.Hostname()

	fmt.Println(string(repDump))
	msg := fmt.Sprintf("Hi from %s (container: %s)\n", *listenAddr, host)
	w.Write([]byte(msg))
}
