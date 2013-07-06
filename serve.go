package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var addr = flag.String("addr", "localhost:10000", "address to serve from")
var verbose = flag.Bool("v", false, "verbose logging")

func main() {
	// flag.Usage = usage // busted usage
	flag.Parse()
	if len(flag.Args()) > 1 {
		log.Fatalf("trailing flags after directory path was given")
	}
	path := flag.Arg(0)

	var err error
	if path == "" {
		path, err = os.Getwd()
	} else {
		path, err = filepath.Abs(path)
	}
	if err != nil {
		log.Fatal(err)
	}

	h := http.FileServer(http.Dir(path))
	if *verbose {
		log.Printf("Serving %s on %s", path, *addr)
		h = &verboseHandler{h}
	}
	log.Fatal(http.ListenAndServe(*addr, h))
}

type verboseHandler struct {
	inner http.Handler
}

func (vh *verboseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Client requested " + r.URL.String())
	vh.inner.ServeHTTP(w, r)
}
