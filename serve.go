package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
)

var addr = flag.String("addr", "localhost:10000", "address to serve from")
var verbose = flag.Bool("v", false, "verbose logging")
var veryVerbose = flag.Bool("vv", false, "very verbose logging, dumping requests")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: serve [DIR]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func main() {
	flag.Usage = usage
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
	if *verbose || *veryVerbose {
		log.Printf("Serving %s on %s", path, *addr)
	}

	if *veryVerbose {
		h = &requestDumpHandler{h}
	} else if *verbose {
		h = &verboseHandler{h}
	}
	log.Fatal(http.ListenAndServe(*addr, h))
}

type verboseHandler struct {
	inner http.Handler
}

func (vh *verboseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s %s", r.Method, r.URL.String(), r.Proto)
	vh.inner.ServeHTTP(w, r)
}

type requestDumpHandler struct {
	inner http.Handler
}

func (rh *requestDumpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	bs, err := httputil.DumpRequest(r, true)
	if err == nil {
		log.Printf("----------------\n%s\n", string(bs))
		log.Println("----------------")
	} else {
		log.Println("Unable to dump request: %s", err)
	}
	rh.inner.ServeHTTP(w, r)
}
