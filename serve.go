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
var quiet = flag.Bool("q", false, "log nothing")
var verbose = flag.Bool("v", false, "verbose logging, dumping requests")

func usage() {
	fmt.Fprintf(os.Stderr, "usage: serve [DIR]\n")
	flag.PrintDefaults()
	os.Exit(1)
}

func die(sfmt string, objs ...interface{}) {
	fmt.Fprintf(os.Stderr, "serve: "+sfmt+"\n", objs...)
	os.Exit(1)
}

func main() {
	flag.Usage = usage
	flag.Parse()
	if len(flag.Args()) > 1 {
		fmt.Fprintf(os.Stderr, "serve: trailing flags were found after directory path\n")
		usage()
	}
	if *quiet && *verbose {
		fmt.Fprintf(os.Stderr, "serve: cannot specify both -q and -v at the same time\n")
		usage()
	}

	path := flag.Arg(0)

	if path == "" {
		cwd, err := os.Getwd()
		if err != nil {
			die("unable to find current working directory")
		}
		path = cwd
	} else {
		absPath, err := filepath.Abs(path)
		if err != nil {
			die("unable to make '%s' an absolute path", path)
		}
		path = absPath
	}

	fi, err := os.Stat(path)
	if err != nil {
		die("'%s' does not exist", path)
	}
	if !fi.IsDir() {
		die("'%s' was not a directory", path)
	}
	h := http.FileServer(http.Dir(path))
	if !*quiet {
		log.Printf("Serving %s on %s", path, *addr)
	}

	if *verbose {
		h = &requestDumpHandler{h}
	} else if !*quiet {
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
